package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microservices-demo/auth/models"
	"net/http"
	"net/url"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mgutz/logxi/v1"
	"gopkg.in/hlandau/passlib.v1"
)

type environment struct {
	logger log.Logger
	db     models.Datastore
}

var env *environment

func main() {
	logger := log.New("")
	logger.Info("auth api starting...")
	db, err := models.NewDB(logger)
	if err != nil {
		logger.Fatal("Could not open database", "err", err)
	}
	env = &environment{logger, db}
	logger.Info("auth api is up and serving on 3001")
	logger.Fatal("Serving", "err", http.ListenAndServe(":3001", handle()))
}

func handle() *mux.Router {
	r := mux.NewRouter()
	// TODO: endpoint to change password (username, old password, new password)
	r.HandleFunc("/auth/token", newToken).Methods("POST")
	r.HandleFunc("/auth/accounts", newAccount).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	return r
}

func notFound(w http.ResponseWriter, r *http.Request) {
	env.logger.Info("Not found: ", "url", r.URL)
	http.Error(w, http.StatusText(404), 404)
}

func newToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	username := r.FormValue("username")
	account, err := env.db.GetAccountByUsername(username)
	if err != nil {
		env.logger.Error("Error getting account by username", "err", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	password := r.FormValue("password")
	newHash, err := passlib.Verify(password, account.PasswordHash)
	if err != nil {
		http.Error(w, http.StatusText(401), 401)
	}
	if newHash != "" {
		env.logger.Info("TODO: save newHash", "newHash", newHash)
	}
	// TODO: POST to Kong and write Kong's response to this response
	data := url.Values{}
	data.Set("client_id", "auth_service")
	data.Set("client_secret", "auth_service_secret")              // TODO: pass in as env var
	data.Set("provision_key", "6ca0c9d2e033476cb57b70b334b524ef") // TODO: pass in as env var
	data.Add("authenticated_userid", strconv.FormatInt(account.Id, 10))
	env.logger.Info("Authenticating user", "newauthenticated_useridHash", strconv.FormatInt(account.Id, 10))
	data.Add("grant_type", "password")
	// TODO: REMOVE THIS testing code and setup proper certs in dev environment
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//client := &http.Client{}
	///////////////////////////////////////////////////////////////////////////
	tokenReq, err := http.NewRequest("POST", "https://kong:8443/api/oauth2/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		env.logger.Error("Error setting up kong oauth2 token request", "err", err, "user_id", account.Id)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tokenReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	tokenReq.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(tokenReq)
	if err != nil {
		env.logger.Error("Error getting token from kong", "err", err, "user_id", account.Id)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		env.logger.Error("Error reading token response from kong", "err", err, "user_id", account.Id)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	body := string(bodyBytes)
	if resp.StatusCode != 200 {
		env.logger.Error("Expected status 200 from kong oauth2/token", "status", resp.StatusCode, "user_id", account.Id, "response", body)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "%s", body)
}

func newAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	username := r.FormValue("username")
	// Check for existing account with that username
	existingAccount, err := env.db.GetAccountByUsername(username)
	if err != nil {
		env.logger.Error("Error getting account by username", "err", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if existingAccount != nil {
		// TODO: informative message that account with that username already exists
		http.Error(w, http.StatusText(422), 422)
		return
	}
	password := r.FormValue("password")
	hash, err := passlib.Hash(password)
	if err != nil {
		env.logger.Error("Error hashing the password", "err", err, "username", username)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	account := models.Account{Username: username, PasswordHash: hash}
	newAccount, err := env.db.CreateAccount(&account)
	if err != nil {
		env.logger.Error("Error creating account", "err", err, "username", username)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	// TODO: return HAL representation of account, not raw json
	j, err := json.MarshalIndent(newAccount, "", "  ")
	if err != nil {
		env.logger.Error("Error serializing Account to JSON", "err", err, "username", username)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "%s", j)
}

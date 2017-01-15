package api

import (
	"dockerized-go-app/rest-api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mgutz/logxi/v1"
	"github.com/nvellon/hal"
)

type environment struct {
	logger log.Logger
	db     models.Datastore
}

var env *environment

// Handlers : given the logger and datastore, returns a pointer to a mux.Router that can handle any HTTP requests
func Handlers(logger log.Logger, db models.Datastore) *mux.Router {
	env = &environment{logger, db}
	r := mux.NewRouter()
	r.HandleFunc("/api", apiRoot).Methods("GET")
	r.HandleFunc("/api/books", env.booksIndex).Methods("GET")
	return r
}

func apiRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	rr := hal.NewResource(models.APIRoot{}, "/api")
	rr.AddNewLink("books", "/api/books")
	// TODO: add links to all possible APIs

	// JSON Encoding
	j, err := json.MarshalIndent(rr, "", "  ")
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, "%s", j)
}

func (env *environment) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := env.db.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rr := hal.NewResource(models.BookResponse{}, "/api/books")
	for _, bk := range bks {
		// Creating HAL Resources
		rb := hal.NewResource(bk, fmt.Sprintf("/api/books/%s", bk.Isbn))
		// TODO: this embedding will create a json object, not an array, if there is only one item here.  Maybe there is
		// some way to always force array type?
		rr.Embed("books", rb)
	}

	// JSON Encoding
	j, err := json.MarshalIndent(rr, "", "  ")
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, "%s", j)
}

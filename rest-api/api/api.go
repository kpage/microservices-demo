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
	r.HandleFunc("/api/orders", env.ordersIndex).Methods("GET")
	return r
}

func apiRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	rr := hal.NewResource(models.APIRoot{}, "/api")
	rr.AddNewLink("restbucks:orders", "/api/orders")
	// TODO: add links to all possible APIs

	// JSON Encoding
	j, err := json.MarshalIndent(rr, "", "  ")
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, "%s", j)
}

func (env *environment) ordersIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	orders, err := env.db.AllOrders()
	if err != nil {
		env.logger.Error("REST API error", "err", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rr := hal.NewResource(models.OrderResponse{}, "/api/orders")
	for _, o := range orders {
		// Creating HAL Resources
		ro := hal.NewResource(o, fmt.Sprintf("/api/orders/%d", o.Id))
		// TODO: this embedding will create a json object, not an array, if there is only one item here.  Maybe there is
		// some way to always force array type?
		rr.Embed("restbucks:orders", ro)
	}

	// JSON Encoding
	j, err := json.MarshalIndent(rr, "", "  ")
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, "%s", j)
}

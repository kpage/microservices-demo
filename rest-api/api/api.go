package api

import (
	"encoding/json"
	"fmt"
	"microservices-demo/rest-api/models"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
func Handlers(logger log.Logger, db models.Datastore) http.Handler {
	env = &environment{logger, db}
	r := mux.NewRouter()
	r.HandleFunc("/api", apiRoot).Methods("GET")
	r.HandleFunc("/api/orders", env.ordersIndex).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	return handlers.LoggingHandler(os.Stdout, r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	env.logger.Info("Not found: ", "url", r.URL)
	http.Error(w, http.StatusText(404), 404)
}

func apiRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	rr := hal.NewResource(models.APIRoot{}, "/api")
	rr.AddNewLink("restbucks:orders", "/api/orders")
	// TODO: add links to all possible APIs

	if err := json.NewEncoder(w).Encode(rr); err != nil {
		fmt.Printf("%s", err)
		return
	}
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
	if err := json.NewEncoder(w).Encode(rr); err != nil {
		fmt.Printf("%s", err)
		return
	}
}

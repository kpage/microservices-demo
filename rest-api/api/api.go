package api

import (
	"dockerized-go-app/rest-api/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mgutz/logxi/v1"
)

type environment struct {
	logger *log.Logger
	db     models.Datastore
}

var env *environment

// Handlers : given the logger and datastore, returns a pointer to a mux.Router that can handle any HTTP requests
func Handlers(logger *log.Logger, db models.Datastore) *mux.Router {
	env = &environment{logger, db}
	r := mux.NewRouter()
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/books", env.booksIndex).Methods("GET")
	//http.Handle("/", r)
	return r
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
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
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

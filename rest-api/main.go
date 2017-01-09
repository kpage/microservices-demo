package main

import (
	"dockerized-go-app/rest-api/models"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mgutz/logxi/v1"
)

type Env struct {
	db models.Datastore
	// TODO: how to pass around this logger to other packages?  Or should each package define its own logger?
	logger log.Logger
}

func main() {
	logger := log.New("")
	logger.Info("rest-api starting...")
	db, err := models.NewDB(logger)
	if err != nil {
		log.Fatal("Could not open database", "err", err)
	}

	env := &Env{db, logger}

	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/books", env.booksIndex)
	http.Handle("/", r)

	logger.Info("rest-api serving on 3001")
	logger.Fatal("Serving", "err", http.ListenAndServe(":3001", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
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

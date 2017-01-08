package main

import (
	"dockerized-go-app/rest-api/models"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB()
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/books", env.booksIndex)
	http.Handle("/", r)

	fmt.Println("Starting up on 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
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

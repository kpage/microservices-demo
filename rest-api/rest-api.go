package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	http.Handle("/", r)
	fmt.Println("Starting up on 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

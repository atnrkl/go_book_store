package main

import (
	"log"
	"net/http"

	"github.com/atnrkl/go_book_store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

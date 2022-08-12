package main

import (
	"aymane/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Routers(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

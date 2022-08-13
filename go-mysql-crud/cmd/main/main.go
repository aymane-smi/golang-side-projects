package main

import (
	"aymane/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Routers(r)
	http.Handle("/", r)
	fmt.Println("server start listening at port : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

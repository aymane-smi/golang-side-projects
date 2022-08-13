package routes

import (
	"aymane/pkg/controllers"

	"github.com/gorilla/mux"
)

var Routers = func(router *mux.Router) {
	router.HandleFunc("/Books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/Books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/Books/{id}", controllers.GetBooksById).Methods("GET")
	router.HandleFunc("/Books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/Books/{id}", controllers.DeleteBook).Methods("DELETE")
}

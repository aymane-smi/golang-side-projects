package routes

import (
	"aymane/pkg/controllers"

	"github.com/gorilla/mux"
)

var routers = func(router *mux.Router) {
	router.HandleFunc("/Books", controllers.getBooks).Methods("GET")
	router.HandleFunc("/Books", controllers.createBook).Methods("POST")
	router.HandleFunc("/Books/{id}", controllers.getBooksById).Methods("GET")
	router.HandleFunc("/Books/{id}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/Books/{id}", controllers.DeleteBook).Methods("DELETE")
}

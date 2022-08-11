package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int64     `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:director`
}

type Director struct {
	fname string `json:"fname"`
	lname string `json:"lname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"message": "all stored movies"}
	response["movies"] = movies
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusAccepted)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Params := mux.Vars(r)
	id, _ := strconv.ParseInt(Params["id"], 0, 8)
	for _, item := range movies {
		if item.ID == id {
			var message string = "movie with id" + Params["id"] + " was found"
			response := map[string]interface{}{"message": message}
			response["movie"] = item
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	var message string = "movie with id " + Params["id"] + " doesn't exist"
	response := map[string]string{"message": message}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func main() {
	movies = append(movies, Movie{ID: 1, Isbn: "jgj45", Title: "test", Director: &Director{"fname", "lname"}})
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

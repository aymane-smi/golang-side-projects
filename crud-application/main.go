package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Director struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type Movie struct {
	ID       int64     `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie
var sequence int64 = 0

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

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Params := mux.Vars(r)
	sequence++
	//tmp := Movie{ID: sequence, Isbn: Params["isbn"], Title: Params["title"], Director: &Director{fname: Params["fname"], lname: Params["lname"]}}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = sequence
	movies = append(movies, movie)
	response := map[string]interface{}{"message": "new movie was add"}
	response["movie"] = movie
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Params := mux.Vars(r)
	id, _ := strconv.ParseInt(Params["id"], 0, 8)
	for index, item := range movies {
		if id == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID
			movies = append(movies, movie)
			w.WriteHeader(http.StatusOK)
			var message string = "movie with id " + Params["id"] + " was updated"
			response := map[string]string{"message": message}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	var message string = "movie with id " + Params["id"] + " doesn't exist"
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Params := mux.Vars(r)
	id, _ := strconv.ParseInt(Params["id"], 0, 8)
	for index, item := range movies {
		if id == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			w.WriteHeader(http.StatusOK)
			var message string = "movie with id " + Params["id"] + "was deleted"
			response := map[string]string{"message": message}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	var message string = "movie with id " + Params["id"] + " doesn't exist"
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)
}

func main() {
	//movies = append(movies, Movie{ID: 1, Isbn: "jgj45", Title: "test", Director: &Director{"fname", "lname"}})
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

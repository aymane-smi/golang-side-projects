package controllers

import (
	"aymane/pkg/models"
	"aymane/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tmp models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	tmp := models.GetAllBooks()
	res, _ := json.Marshal(tmp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	Params := mux.Vars(r)
	id, err := strconv.ParseInt(Params["id"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	Params := mux.Vars(r)
	id, err := strconv.ParseInt(Params["id"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book, db := models.GetBookById(id)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	Params := mux.Vars(r)
	id, err := strconv.ParseInt(Params["id"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	tmp = models.DeleteBookById(id)
	res, _ := json.Marshal(tmp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/models"
	"example.com/services"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookData models.BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateBook(bookData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func ReadBooks(w http.ResponseWriter, r *http.Request) {
	result, err := services.ReadBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_slug := params["book_slug"]

	book, err := services.ReadBook(book_slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_slug := params["book_slug"]

	var bookData models.BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := services.UpdateBook(book_slug, bookData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_slug := params["book_slug"]

	if err := services.DeleteBook(book_slug); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
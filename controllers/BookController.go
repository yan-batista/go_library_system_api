package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/models"
	"example.com/services"
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
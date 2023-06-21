package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/models"
	"example.com/services"
)

func RentBook(w http.ResponseWriter, r *http.Request) {
	var rentData models.Rent
	if err := json.NewDecoder(r.Body).Decode(&rentData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.RentBook(rentData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	var rentData models.Rent
	if err := json.NewDecoder(r.Body).Decode(&rentData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return
	}

	if err := services.ReturnBook(rentData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func ExtendRent(w http.ResponseWriter, r *http.Request) {
	var rentData models.Rent
	if err := json.NewDecoder(r.Body).Decode(&rentData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.ExtendRent(rentData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func FindRentedBooks(w http.ResponseWriter, r *http.Request) {
	// query

	result, err := services.FindRentedBooks(); 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
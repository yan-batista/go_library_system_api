package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/models"
	"example.com/services"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func FindUserByQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
    first_name := query.Get("first_name")
    last_name := query.Get("last_name")
    email := query.Get("email")

	result, err := services.FindUserByQuery(first_name, last_name, email); 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]

	result, err := services.FindUserByQuery("", "", email);
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(result) == 0 {
		http.Error(w, "usuário não encontrado", http.StatusNotFound)
		return 
	}

	if err := services.DeleteUser(email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
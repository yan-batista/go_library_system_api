package routes

import (
	"example.com/controllers"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	bookRoutes := r.PathPrefix("/books").Subrouter()
	//userRoutes := r.PathPrefix("/users").Subrouter()

	bookRoutes.HandleFunc("/", controllers.CreateBook).Methods("POST")
	bookRoutes.HandleFunc("/", controllers.ReadBooks).Methods("GET")
	//bookRoutes.HandleFunc("/{id}", GetBook).Methods("GET")
	//bookRoutes.HandleFunc("/{id}", UpdateBook).Methods("PUT")
	//bookRoutes.HandleFunc("/{id}", DeleteBook).Methods("DELETE")
	//bookRoutes.HandleFunc("", GetBookByQuery).Methods("GET")

	return r
}
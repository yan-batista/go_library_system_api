package routes

import (
	"example.com/controllers"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	bookRoutes := r.PathPrefix("/books").Subrouter()
	userRoutes := r.PathPrefix("/users").Subrouter()

	bookRoutes.HandleFunc("", controllers.CreateBook).Methods("POST")
	bookRoutes.HandleFunc("", controllers.FindByQuery).Methods("GET")
	bookRoutes.HandleFunc("/{book_slug}", controllers.ReadBook).Methods("GET")
	bookRoutes.HandleFunc("/{book_slug}", controllers.UpdateBook).Methods("PUT")
	bookRoutes.HandleFunc("/{book_slug}", controllers.DeleteBook).Methods("DELETE")

	userRoutes.HandleFunc("", controllers.CreateUser).Methods("POST")
	userRoutes.HandleFunc("", controllers.FindUserByQuery).Methods("GET")
	userRoutes.HandleFunc("/{email}", controllers.DeleteUser).Methods("DELETE")

	return r
}
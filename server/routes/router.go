package routes

import (
	"example.com/controllers"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	bookRoutes := r.PathPrefix("/api/books").Subrouter()
	userRoutes := r.PathPrefix("/api/users").Subrouter()
	rentRoutes := r.PathPrefix("/api/rent").Subrouter()

	bookRoutes.HandleFunc("", controllers.CreateBook).Methods("POST")
	bookRoutes.HandleFunc("", controllers.FindByQuery).Methods("GET")
	bookRoutes.HandleFunc("/{book_slug}", controllers.ReadBook).Methods("GET")
	bookRoutes.HandleFunc("/{book_slug}", controllers.UpdateBook).Methods("PUT")
	bookRoutes.HandleFunc("/{book_slug}", controllers.DeleteBook).Methods("DELETE")

	userRoutes.HandleFunc("", controllers.CreateUser).Methods("POST")
	userRoutes.HandleFunc("", controllers.FindUserByQuery).Methods("GET")
	userRoutes.HandleFunc("/{email}", controllers.UpdateUser).Methods("PUT")
	userRoutes.HandleFunc("/{email}", controllers.DeleteUser).Methods("DELETE")

	rentRoutes.HandleFunc("", controllers.RentBook).Methods("POST")
	rentRoutes.HandleFunc("", controllers.FindRentedBooks).Methods("GET")
	rentRoutes.HandleFunc("", controllers.ExtendRent).Methods("PUT")
	rentRoutes.HandleFunc("", controllers.ReturnBook).Methods("DELETE")
	
	return r
}

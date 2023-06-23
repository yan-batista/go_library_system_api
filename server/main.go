package main

import (
	"net/http"

	"example.com/db/seed"
	"example.com/routes"
)

func main() {
	// create a router and server
	r := routes.CreateRouter()

	seed.SeedDB()	

	http.ListenAndServe(":3000", r)
}
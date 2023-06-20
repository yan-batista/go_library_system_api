package models

import (
	"time"
)

type Book struct {
	Id        	int       	`json:"id"`
	Name      	string    	`json:"name"`
	Slug      	string    	`json:"slug"`
	Author    	string    	`json:"author"`
	Publisher 	string    	`json:"publisher"`
	ISBN      	string    	`json:"isbn"`
	Quantity  	int       	`json:"quantity"`
	Description string 		`json:"description"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

type BookDTO struct {
	Name      	string    	`json:"name"`
	Author    	string    	`json:"author"`
	Publisher 	string    	`json:"publisher"`
	ISBN      	string    	`json:"isbn"`
	Quantity  	int       	`json:"quantity"`
	Description string 		`json:"description"`
}
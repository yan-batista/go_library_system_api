package models

import "time"

type Rent struct {
	BookSlug   string    `json:"book_slug"`
	UserEmail  string    `json:"user_email"`
	ReturnDate time.Time `json:"return_date"`
}

type RentedBook struct {
	Name      	string    	`json:"name"`
	Slug      	string    	`json:"slug"`
	Author    	string    	`json:"author"`
	Publisher 	string    	`json:"publisher"`
	ISBN      	string    	`json:"isbn"`
	Email  		string    	`json:"email"`
	Phone  		string    	`json:"Phone"`
	ReturnDate 	time.Time 	`json:"return_date"`
}
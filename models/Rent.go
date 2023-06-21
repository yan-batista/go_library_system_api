package models

import "time"

type Rent struct {
	BookSlug   string    `json:"book_slug"`
	UserEmail  string    `json:"user_email"`
	ReturnDate time.Time `json:"return_date"`
}

type RentedBook struct {
	BookName    string    	`json:"book_name"`
	Slug      	string    	`json:"slug"`
	Author    	string    	`json:"author"`
	Publisher 	string    	`json:"publisher"`
	ISBN      	string    	`json:"isbn"`
	UserName    string    	`json:"user_name"`
	Email  		string    	`json:"email"`
	Phone  		string    	`json:"Phone"`
	ReturnDate 	time.Time 	`json:"return_date"`
}
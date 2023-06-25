package models

import "time"

type User struct {
	Id        	int       	`json:"id"`
	FirstName  	string    	`json:"first_name"`
	LastName  	string    	`json:"last_name"`
	Email  		string    	`json:"email"`
	Phone  		string    	`json:"Phone"`
	Debt      	int       	`json:"debt"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
}

type UserDTO struct {
	FirstName  	string    	`json:"first_name"`
	LastName  	string    	`json:"last_name"`
	Email  		string    	`json:"email"`
	Phone  		string    	`json:"Phone"`
	Debt      	int       	`json:"debt"`
}
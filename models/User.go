package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  int       `json:"username"`
	Password  int       `json:"password"`
	Debt      int       `json:"debt"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
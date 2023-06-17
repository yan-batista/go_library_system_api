# Library System API

This project is a RESTful API for a library system, written in Go. It allows users to manage books and their information. The API also provides functionality for managing users and borrowing/returning books.

## API Documentation

## Env. Variables

Necessary environment variables to run the project.

`CONNECTION_STRING`: Connection string for the database

## Functionalities

- Create, Read, Update and Delete books
- Find books by Name, Author or ISBN
- Create, Read, Update and Delete Users
- Rent book
- Return Book
- Get user's current books

## Models

Book

```go
type Book struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	ISBN      string    `json:"isbn"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

User

```go
type User struct {
	Id        int       `json:"id"`
	Username  int       `json:"username"`
	Password  int       `json:"password"`
	Debt      int       `json:"debt"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

## Use Cases

- It's not possible to create a book with a duplicate ISBN or Slug
- User can't rent more than 5 books
- User can't rent a book out of stock
- User can't rent multiple copies of the same book
- User can only return a book in their possession
- User can rent a book for 7, 15 or 30 days
- In case a user doesn't return a book, they must pay $3 per day after return date. On return, the total fee is calculated.

## What I've learned

# Library System API

This project is a RESTful API for a library system, written in Go. It allows users to manage books and their information. The API also provides functionality for managing users and renting and returning books.

# Env. Variables

Necessary environment variables to run the project.

`CONNECTION_STRING`: Connection string for the database

# Functionalities

- Create, Read, Update and Delete books
- Find books by Name, Author or ISBN
- Create, Read, Update and Delete Users
- Rent book
- Return Book
- Get user's current books

# Use Cases

- It's not possible to create a book with a duplicate ISBN or Slug
- User can't rent more than 5 books
- User can't rent a book out of stock
- User can't rent multiple copies of the same book
- User can only return a book in their possession
- User can rent a book for 30 days max
- In case a user doesn't return a book, they must pay $3 per day after return date. On return, the total fee is calculated.

# What I've learned

I built this project with the purpose of learning Golang, so there's a bunch to unpack in this section.

## Creating a server

Golang comes with a package that provides all the functionality for creating an HTTP client or server implementation.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.ListenAndServe(":8080", nil)
}
```

## Handling requests

Requests can be handled using a handler function, which is a function that takes two arguments: a `http.ResponseWriter` and a `http.Request`. The http.ResponseWriter is used to write the response data and headers, while the http.Request contains information about the request such as the method, URL, headers, body, etc.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //...
})
```

## Routing

While the net/http package provides a basic functionality for creating a web server and handling routes, I'm using the gorilla/mux package for the routing part. Gorilla/mux is a powerful HTTP router and URL matcher.

To use gorilla/mux, I first had to install it using the command `go get -u github.com/gorilla/mux`. Then I imported it in a separate `router.go` file.

I then create the router, and subrouters for books, users and rent (relationship between books and users).

```go
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
    //...

	userRoutes.HandleFunc("", controllers.CreateUser).Methods("POST")
    //...

	rentRoutes.HandleFunc("", controllers.RentBook).Methods("POST")
    //...

	return r
}
```

### Queries and Params

If you want to use queries on the routes, you make a route normally, and get the queries from the URL in the controller from the `http.Request`.

```go
// BookController.go
// route -> /books

func FindByQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
    name := query.Get("name")
    author := query.Get("author")
    publisher := query.Get("publisher")
    isbn := query.Get("isbn")
}
```

If you want to use params instead of queries, you can use `mux.Vars` function. The `mux.Vars` takes a http.Request as an argument and returns a map with both key and value string that represents the params. The keys are the names of the params and the values are the extracted data from the request URL.

```go
// BookController.go
// route -> /books/{book_slug}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book_slug := params["book_slug"]
}
```

## Connecting to a database

For this project I used a mySQL database, using the go-sql-driver/mysql. To install the driver, we can use the command `go get -u github.com/go-sql-driver/mysql`. Then we can import it in our code using the blank identifier \_:

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
```

To open a connection to the MySQL database, we need to use the sql.Open function. This function takes two arguments: the driver name and the data source name (DSN). The driver name is simply "mysql" for this driver. The DSN is a string that contains the connection information for the database, such as the host, port, user, password, database name, etc.

```go
db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/mydb?parseTime=true")
if err != nil {
    // handle error
}
```

Instead of hardcoding the connection string for the database, we should hide it with environment variables.

## Using environment variables

To load environment variables from a .env file, we can use the joho/godotenv package. This package provides functions for reading and writing .env files and loading them into our code.

To install the package, we can use the command `go get -u github.com/joho/godotenv`. Then we can import it in our code.

To load environment variables from a .env file, we can use the `godotenv.Load` function. This function takes an optional argument of one or more file names. If no argument is given, it will look for a file named .env in the current directory.

```go
package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	connection_string := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("mysql", connection_string)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
```

## Database functions

To execute statements that **modify** data in the database, such as **INSERT, UPDATE or DELETE**, we can use the `Exec` function of the `sql.DB` object. This function takes a SQL statement as an argument and returns a `sql.Result` object and an error. The sql.Result object contains information about the execution result, such as the number of affected rows or the last inserted ID.

```go
func ExtendRent(rentData models.Rent) error {
	query := `UPDATE rent SET return_date = ? WHERE book_slug = ? AND user_email = ?`
	if _, err := connection.Exec(query, rentData.ReturnDate, rentData.BookSlug, rentData.UserEmail); err != nil {
		return err
	}
	return nil
}
```

To execute statements that **return rows** from the database, such as **SELECT**, we can use the `Query` or `QueryRow` functions of the `sql.DB` object. These functions take a SQL statement as an argument and return a `sql.Rows` or a `sql.Row` object and an error. The `sql.Rows` object represents a cursor that allows us to iterate over the rows returned by the query. The `sql.Row` object represents a single row returned by the query.

```go
func FindRentedBooks(book_slug, user_email string) ([]models.RentedBook, error) {
	query := `...`
    // exec the query
	rows, err := connection.Query(query, book_slug, book_slug, user_email, user_email);
	if err != nil {
		// handle error
	}
    // defer closes the rows object
	defer rows.Close()

	var results []models.RentedBook
    // iterate over the rows
	for rows.Next() {
		var result models.RentedBook

        // scan each row into variables
		err := rows.Scan(&result.BookName, &result.Slug, &result.Author, &result.Publisher, &result.ISBN, &result.UserName ,&result.Email, &result.Phone, &result.ReturnDate);
		if err != nil {
			// handle error
		}

        // put all scanned rows into a slice
		results = append(results, result)
	}

    // return data
	return results, nil
}
```

```go
func ReadBook(book_slug string) (models.Book, error) {
	query = "SELECT * from books WHERE slug = ?"

	var book models.Book

    // queries row, and scan the row into variables
	if err := connection.QueryRow(query, book_slug).Scan(&book.Id, &book.Name, &book.Slug, &book.Author, &book.Publisher, &book.ISBN, &book.Quantity, &book.Description, &book.CreatedAt, &book.UpdatedAt); err != nil {
		// handle error
	}

    // return scanned data
	return book, nil
}
```

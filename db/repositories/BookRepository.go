package repositories

import (
	"example.com/db"
	"example.com/models"
)

var query string
var connection = db.CreateConnection()

func CreateBook(book models.Book) error {
	query = "INSERT INTO books (name, slug, author, publisher, isbn, quantity, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := connection.Exec(query, book.Name, book.Slug, book.Author, book.Publisher, book.ISBN, book.Quantity, book.CreatedAt, book.UpdatedAt)
	if err != nil {
		return err;
	}
	
	return nil
}

func ReadBooks() ([]models.Book, error) {
	query = "SELECT * FROM books"

	rows, err := connection.Query(query);
	if err != nil {
		return nil, err;
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book

		err := rows.Scan(&book.Id, &book.Name, &book.Slug, &book.Author, &book.Publisher, &book.ISBN, &book.Quantity, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
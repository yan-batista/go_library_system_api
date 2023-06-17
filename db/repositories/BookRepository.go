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

func ReadBook(book_slug string) (models.Book, error) {
	query = "SELECT * from books WHERE slug = ?"

	var book models.Book

	if err := connection.QueryRow(query, book_slug).Scan(&book.Id, &book.Name, &book.Slug, &book.Author, &book.Publisher, &book.ISBN, &book.Quantity, &book.CreatedAt, &book.UpdatedAt); err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func UpdateBook(book models.Book) error {
	query = "UPDATE books SET name = ?, slug = ?, author = ?, publisher = ?, isbn = ?, quantity = ?, updated_at = ? WHERE id = ?"
	if _, err := connection.Exec(query, book.Name, book.Slug, book.Author, book.Publisher, book.ISBN, book.Quantity, book.UpdatedAt, book.Id); err != nil {
		return err
	}
	return nil
}

func DeleteBook(book_slug string) error {
	query = "DELETE FROM books WHERE slug = ?"
	if _, err := connection.Exec(query, book_slug); err != nil {
		return err
	}
	return nil
}
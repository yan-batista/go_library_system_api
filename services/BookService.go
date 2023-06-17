package services

import (
	"errors"
	"time"

	"example.com/db/repositories"
	"example.com/models"
	"example.com/utils"
)

func CreateBook(bookData models.BookDTO) error {
	// bookData fields cannot be empty
	if (bookData.Name == "" || bookData.Author == "" || bookData.Publisher == "" || bookData.ISBN == "") {
		return errors.New("all fields are obligatory")
	}

	if (bookData.Quantity < 0) {
		return errors.New("invalid value for quantity")
	}

	book_slug := utils.CreateSlug(bookData.Name)

	book := models.Book{
		Id: 0,
		Name: bookData.Name,
		Slug: book_slug,
		Author: bookData.Author,
		Publisher: bookData.Publisher,
		ISBN: bookData.ISBN,
		Quantity: bookData.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repositories.CreateBook(book); err != nil {
		return err;
	}

	return nil
}

func ReadBooks() ([]models.Book, error) {
	result, err := repositories.ReadBooks()
	if err != nil {
		return nil, err
	}

	return result, nil
}
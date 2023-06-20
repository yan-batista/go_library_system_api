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
	
	book_slug := utils.CreateSlug(bookData.Name, bookData.Author)

	// if book_slug is registered, return error
	if _, err := repositories.ReadBook(book_slug); err == nil {
		return errors.New("book already exists")
	}
	// if book_isbn is registered, return error
	result, err := repositories.FindByQuery("", "", "", bookData.ISBN); 
	if err != nil {
		return err
	}
	if len(result) != 0 {
		return errors.New("book with the same ISBN already exists")
	}

	book := models.Book{
		Id: 0,
		Name: bookData.Name,
		Slug: book_slug,
		Author: bookData.Author,
		Publisher: bookData.Publisher,
		ISBN: bookData.ISBN,
		Quantity: bookData.Quantity,
		Description: bookData.Description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repositories.CreateBook(book); err != nil {
		return err;
	}

	return nil
}

func FindByQuery(name, author, publisher, isbn string) ([]models.Book, error) {
	result, err := repositories.FindByQuery(name, author, publisher, isbn)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ReadBook(book_slug string) (models.Book, error) {
	result, err := repositories.ReadBook(book_slug)
	if err != nil {
		return models.Book{}, err
	}
	return result, nil
}

func UpdateBook(book_slug string, bookData models.BookDTO) error {
	// find book
	result, err := repositories.ReadBook(book_slug); 
	if err != nil {
		return err
	}

	name := bookData.Name
	if bookData.Name == "" {
		name = result.Name
	}
	author := bookData.Author
	if bookData.Author == "" {
		author = result.Author
	}
	publisher := bookData.Publisher
	if bookData.Publisher == "" {
		publisher = result.Publisher
	}
	isbn := bookData.ISBN
	if bookData.ISBN == "" {
		isbn = result.ISBN
	}
	quantity := bookData.Quantity
	if quantity < 0 {
		quantity = result.Quantity
	}
	description := bookData.Description
	if description == "" {
		description = result.Description
	}

	// update model data
	book := models.Book{
		Id: result.Id,
		Name: name,
		Slug: utils.CreateSlug(name, author),
		Author: author,
		Publisher: publisher,
		ISBN: isbn,
		Description: description,
		Quantity: quantity,
		CreatedAt: result.CreatedAt,
		UpdatedAt: time.Now(),
	}

	// update repo
	if err := repositories.UpdateBook(book); err != nil {
		return err
	}

	return nil
}

func DeleteBook(book_slug string) error {
	// Verifica se a slug existe, antes de deletar
	if _, err := repositories.ReadBook(book_slug); err != nil {
		return err
	}

	if err := repositories.DeleteBook(book_slug); err != nil {
		return err
	}
	return nil
}
package services

import (
	"errors"
	"time"

	"example.com/db/repositories"
	"example.com/models"
)

func RentBook(rentData models.Rent) error {
	// data can't be empty
	if rentData.BookSlug == "" || rentData.UserEmail == "" {
		return errors.New("data can't be empty")
	}

	// find book
	if result, err := ReadBook(rentData.BookSlug); err != nil {
		return err
	} else if result.Quantity < 0 {
		return errors.New("there are no copies available")
	}

	// find user
	if result, err := FindUserByQuery("", "", rentData.UserEmail); err != nil {
		return err
	} else if len(result) == 0 {
		return errors.New("user not found")
	}

	// check if user already rent that book

	if err := validatesDate(rentData.ReturnDate); err != nil {
		return err
	}

	if err := repositories.RentBook(rentData); err != nil {
		return err
	}

	// update quantity -1

	return nil
}

func ReturnBook(rentData models.Rent) error {
	if rentData.BookSlug == "" || rentData.UserEmail == "" {
		return errors.New("data can't be empty")
	}

	// find book to see if it is are valid
	if _, err := ReadBook(rentData.BookSlug); err != nil {
		return err
	}

	// find user to see if it is are valid
	if result, err := FindUserByQuery("", "", rentData.UserEmail); err != nil {
		return err
	} else if len(result) == 0 {
		return errors.New("user not found")
	}

	// check if user actually rented book

	// check if user should pay fees. Check if book is late

	if err := repositories.ReturnBook(rentData); err != nil {
		return err
	}

	// update quantity +1
	
	return nil
}

func ExtendRent(rentData models.Rent) error {
	// data can't be empty
	if rentData.BookSlug == "" || rentData.UserEmail == "" {
		return errors.New("data can't be empty")
	}

	// find Rent

	if err := validatesDate(rentData.ReturnDate); err != nil {
		return err
	}

	if err := repositories.ExtendRent(rentData); err != nil {
		return err
	}

	return nil
}

func FindRentedBooks() ([]models.RentedBook, error) {
	result, err := repositories.FindRentedBooks(); 
	if err != nil {
		return nil, err
	}

	return result, nil
}

func validatesDate(date time.Time) error {
	if date.Before(time.Now()) {
		return errors.New("date can't be before current date")
	}

	timeLeft := time.Until(date)
	daysLeft := int(timeLeft.Hours() / 24)

	if daysLeft > 30 {
		return errors.New("user can only rent a book for a maximum of 30 days")
	}

	return nil
}
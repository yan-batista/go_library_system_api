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
	book_result, err := ReadBook(rentData.BookSlug); 
	if err != nil {
		return err
	} else if book_result.Quantity <= 0 {
		return errors.New("there are no copies available")
	}

	// find user
	if result, err := FindUserByQuery("", "", rentData.UserEmail); err != nil {
		return err
	} else if len(result) == 0 {
		return errors.New("user not found")
	}

	// check if user already rented that book
	result, err := FindRentedBooks(rentData.BookSlug, rentData.UserEmail)
	if err != nil {
		return err
	}
	if len(result) > 0 {
		return errors.New("user already rented this book")
	}

	// limit the number of books rented
	maxRent := 5
	if result, _ := FindRentedBooks("", rentData.UserEmail); len(result) >= maxRent {
		return errors.New("user can't rent more than 5 books")
	}

	// validates date
	if err := validatesDate(rentData.ReturnDate); err != nil {
		return err
	}

	if err := repositories.RentBook(rentData); err != nil {
		return err
	}

	// update quantity -1
	if err := UpdateBook(rentData.BookSlug, models.BookDTO{Name: "", Author: "", Publisher: "", ISBN: "", Quantity: book_result.Quantity - 1, Description: ""}); err != nil {
		return err
	}

	return nil
}

func ReturnBook(rentData models.Rent) error {
	if rentData.BookSlug == "" || rentData.UserEmail == "" {
		return errors.New("data can't be empty")
	}

	// find book to see if it is valid
	book_result, err := ReadBook(rentData.BookSlug); 
	if err != nil {
		return err
	}

	// find user to see if it is are valid
	if result, err := FindUserByQuery("", "", rentData.UserEmail); err != nil {
		return err
	} else if len(result) == 0 {
		return errors.New("user not found")
	}

	// check if user actually rented book
	result, err := FindRentedBooks(rentData.BookSlug, rentData.UserEmail)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("book not rented by user")
	}

	// check if user should pay fees. Check if book is late
	fee := calculateFee(result[0].Users[0].ReturnDate)
	if fee > 0 {
		if err := UpdateUser(rentData.UserEmail, models.UserDTO{FirstName: "", LastName: "", Email: "", Phone: "", Debt: fee}); err != nil {
			return err
		}
	}

	if err := repositories.ReturnBook(rentData); err != nil {
		return err
	}

	if err := UpdateBook(result[0].Slug, models.BookDTO{Name: "", Author: "", Publisher: "", ISBN: "", Quantity: book_result.Quantity + 1, Description: ""}); err != nil {
		return err
	}
	
	return nil
}

func ExtendRent(rentData models.Rent) error {
	// data can't be empty
	if rentData.BookSlug == "" || rentData.UserEmail == "" {
		return errors.New("data can't be empty")
	}

	result, err := FindRentedBooks(rentData.BookSlug, rentData.UserEmail)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("book not rented by user")
	}

	if err := validatesDate(rentData.ReturnDate); err != nil {
		return err
	}

	if err := repositories.ExtendRent(rentData); err != nil {
		return err
	}

	return nil
}

func FindRentedBooks(book_slug, user_email string) ([]models.RentedBookInfo, error) {
	result, err := repositories.FindRentedBooks(book_slug, user_email); 
	if err != nil {
		return nil, err
	}

	// change data format
	var formattedArray []models.RentedBookInfo
	if len(result) > 0 {
		for _, rented := range result {
			if !checkFormattedContains(rented.Slug, &formattedArray) {
				usersInfo := searchUsersInfo(rented.Slug, &result)
				bookInfo := models.RentedBookInfo{
					BookName: rented.BookName,
					Slug: rented.Slug,
					Author: rented.Author,
					Publisher: rented.Publisher,
					ISBN: rented.ISBN,
					Users: usersInfo,
				}
				formattedArray = append(formattedArray, bookInfo)
			}
		}
	}

	return formattedArray, nil
}

// check if book slug is already formatted
func checkFormattedContains(slug string, formattedArray *[]models.RentedBookInfo) bool {
	for _, data := range *formattedArray {
		if data.Slug == slug {
			return true
		}
	}
	return false
}

// search all users that rent a book by its slug
func searchUsersInfo(slug string, rentedBookArray *[]models.RentedBook) []models.UserInfo {
	var userArray []models.UserInfo

	for _, data := range *rentedBookArray {
		if data.Slug == slug {
			userInfo := models.UserInfo{
				UserName: data.UserName,
				Email: data.Email,
				Phone: data.Phone,
				ReturnDate: data.ReturnDate,
			}
			userArray = append(userArray, userInfo)
		}
	}

	return userArray
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

func calculateFee(return_date time.Time) int {
	money_per_day := 3

	duration := time.Since(return_date)
	days_late := int(duration.Hours() / 24)

	fee := 0
	if days_late > 0 {
		fee = days_late * money_per_day
	}
	return fee
}
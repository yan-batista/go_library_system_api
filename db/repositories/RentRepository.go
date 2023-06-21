package repositories

import (
	"example.com/models"
)

func RentBook(rentData models.Rent) error {
	query = `INSERT INTO rent (book_slug, user_email, return_date) VALUES (?, ?, ?)`
	if _, err := connection.Exec(query, rentData.BookSlug, rentData.UserEmail, rentData.ReturnDate); err != nil {
		return err
	}
	return nil
}

func ReturnBook(rentData models.Rent) error {
	query = `DELETE FROM rent where book_slug = ? AND user_email = ?`
	if _, err := connection.Exec(query, rentData.BookSlug, rentData.UserEmail); err != nil {
		return err
	}
	return nil
}

func ExtendRent(rentData models.Rent) error {
	query := `UPDATE rent SET return_date = ? WHERE book_slug = ? AND user_email = ?`
	if _, err := connection.Exec(query, rentData.ReturnDate, rentData.BookSlug, rentData.UserEmail); err != nil {
		return err
	}
	return nil
}

func FindRentedBooks() ([]models.RentedBook, error) {
	query := `
		SELECT books.name, books.slug, books.author, books.publisher, books.isbn, users.email, users.phone, rent.return_date
		FROM users 
		INNER JOIN rent
		ON rent.user_email = users.email
		INNER JOIN books
		ON books.slug = rent.book_slug
		ORDER BY books.slug
	`
	rows, err := connection.Query(query); 
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.RentedBook
	for rows.Next() {
		var result models.RentedBook

		err := rows.Scan(&result.Name, &result.Slug, &result.Author, &result.Publisher, &result.ISBN, &result.Email, &result.Phone, &result.ReturnDate);
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}


package repositories

import (
	"example.com/models"
)

func CreateUser(user models.UserDTO) error {
	query = `INSERT INTO users (first_name, last_name, email, phone, debt) VALUES (?, ?, ?, ?, ?)`

	if _, err := connection.Exec(query, user.FirstName, user.LastName, user.Email, user.Phone, user.Debt); err != nil {
		return err
	}

	return nil
}

func FindUserByQuery(first_name, last_name, email string) ([]models.User, error) {
	query = `SELECT * FROM users WHERE (? = "" OR first_name LIKE CONCAT('%', ?, '%')) AND (? = "" OR last_name LIKE CONCAT('%', ?, '%')) AND (? = "" OR email LIKE CONCAT('%', ?, '%'))`

	rows, err := connection.Query(query, first_name, first_name, last_name, last_name, email, email); 
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Debt, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func DeleteUser(email string) error {
	query = `DELETE FROM users WHERE email = ?`
	if _, err := connection.Exec(query, email); err != nil {
		return err
	}
	return nil
}
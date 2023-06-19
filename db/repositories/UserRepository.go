package repositories

import (
	"errors"
	"fmt"

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
	query = `SELECT * FROM users WHERE (? = "" OR first_name LIKE CONCAT('%', ?, '%')) AND (? = "" OR last_name LIKE CONCAT('%', ?, '%')) AND (? = "" OR email = ?)`

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

func UpdateUser(user_email string, user_data models.UserDTO) error {
    setClause := ""
    args := []interface{}{}
    if user_data.FirstName != "" {
        setClause += "first_name = ?"
        args = append(args, user_data.FirstName)
    }
    if user_data.LastName != "" {
        if setClause != "" {
            setClause += ", "
        }
        setClause += "last_name = ?"
        args = append(args, user_data.LastName)
    }
    if user_data.Email != "" {
        if setClause != "" {
            setClause += ", "
        }
        setClause += "email = ?"
        args = append(args, user_data.Email)
    }
    if user_data.Phone != "" {
        if setClause != "" {
            setClause += ", "
        }
        setClause += "phone = ?"
        args = append(args, user_data.Phone)
    }
    if user_data.Debt >= 0 {
        if setClause != "" {
            setClause += ", "
        }
        setClause += "debt = ?"
        args = append(args, user_data.Debt)
    }

    if setClause == "" {
        return errors.New("no fields to update")
    }

    query := fmt.Sprintf("UPDATE users SET %s WHERE email = ?", setClause)
    args = append(args, user_email)
    _, err := connection.Exec(query, args...)
    return err
}

func DeleteUser(email string) error {
	query = `DELETE FROM users WHERE email = ?`
	if _, err := connection.Exec(query, email); err != nil {
		return err
	}
	return nil
}
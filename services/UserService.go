package services

import (
	"errors"
	"regexp"

	"example.com/db/repositories"
	"example.com/models"
)

func CreateUser(user models.UserDTO) error {
	// check if email is valid
    if(!isEmailValid(user.Email)) {
		return errors.New("invalid email")
	} 

	// check if email is registred
	result, err := FindUserByQuery("", "", user.Email); 
	if err != nil {
		return err;
	}
	if len(result) > 0 {
		return errors.New("email already registred")
	}

	// create user
	if err := repositories.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func FindUserByQuery(first_name, last_name, email string) ([]models.User, error) {
	if(email != "" && !isEmailValid(email)) {
		return nil, errors.New("invalid email")
	}

	result, err := repositories.FindUserByQuery(first_name, last_name, email); 
	if err != nil {
		return nil, err;
	}

	return result, nil
}

func DeleteUser(email string) error {
	if(!isEmailValid(email)) {
		return errors.New("invalid email")
	}

	if err := repositories.DeleteUser(email); err != nil {
		return err
	}
	
	return nil
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&’*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
    isValid := emailRegex.MatchString(email)
	return isValid
}
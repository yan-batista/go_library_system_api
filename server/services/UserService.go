package services

import (
	"errors"

	"example.com/db/repositories"
	"example.com/models"
	"example.com/utils"
)

func CreateUser(user models.UserDTO) error {
	// check if email is valid
    if(!utils.IsEmailValid(user.Email)) {
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
	if(email != "" && !utils.IsEmailValid(email)) {
		return nil, errors.New("invalid email")
	}

	result, err := repositories.FindUserByQuery(first_name, last_name, email); 
	if err != nil {
		return nil, err;
	}

	return result, nil
}

func UpdateUser(user_email string, user_data models.UserDTO) error {
	// search user based on email on params
	result, err := FindUserByQuery("", "", user_email)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("user not found")
	}

	// if we have an email to update
	if user_data.Email != "" {
		// check if new email is valid
		if(!utils.IsEmailValid(user_data.Email)) {
			return errors.New("invalid email")
		}

		// check if new email is not in use
		result, err := FindUserByQuery("", "", user_data.Email)
		if err != nil {
			return err
		}
		if len(result) > 0 {
			return errors.New("email already in use")
		}
	}

	err = repositories.UpdateUser(user_email, user_data);
	return err
}

func DeleteUser(email string) error {
	if(!utils.IsEmailValid(email)) {
		return errors.New("invalid email")
	}

	if err := repositories.DeleteUser(email); err != nil {
		return err
	}
	
	return nil
}


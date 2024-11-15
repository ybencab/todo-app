package validators

import (
	"errors"
	"regexp"
) 

func ValidateRegisterUserRequest(username, email, password string) error {
	if err := validateUsername(username); err != nil {
		return err
	}
	if err := validateEmail(email); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}
	return nil
}

func validateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	// TODO: Add more password validation rules
	return nil
}

func validateEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

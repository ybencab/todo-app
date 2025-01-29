package validators

import "errors"

func ValidateCreateTodoRequest(title, description string) error {
	if title == "" {
		return errors.New("title is required")
	}
	if description == "" {
		return errors.New("description is required")
	}
	return nil
}

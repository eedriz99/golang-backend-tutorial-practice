package validators

import (
	"BackendDevelopment/tutorial/models"
	"errors"
)

func IsValidUser(u models.User) error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	} else if u.LastName == "" {
		return errors.New("last name is required")
	} else if u.Email == "" {
		return errors.New("email is required")
	} else {
		return nil
	}
}

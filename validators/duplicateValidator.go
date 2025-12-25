package validators

import (
	"BackendDevelopment/tutorial/models"
	"errors"
)

func IsDuplicateUser(u models.User, users []models.User) error {
	for _, user := range users {
		if u.Email == user.Email {
			return errors.New("email address has been taken")
		}
	}
	return nil
}

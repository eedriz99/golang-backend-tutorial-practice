package validators

import "BackendDevelopment/tutorial/models"

func IsDuplicateUser(u models.User, users []models.User) (bool, string) {
	for _, user := range users {
		if u.Email == user.Email {
			return true, "Email address has been taken!"
		}
	}
	return false, ""
}

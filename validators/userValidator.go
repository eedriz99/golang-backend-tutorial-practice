package validators

import "BackendDevelopment/tutorial/models"

func IsValidUser(u models.User) (bool, string) {
	if u.FirstName == "" {
		return false, "First name is required"
	} else if u.LastName == "" {
		return false, "Last name is required"
	} else if u.Email == "" {
		return false, "Email is required"
	} else {
		return true, ""
	}
}

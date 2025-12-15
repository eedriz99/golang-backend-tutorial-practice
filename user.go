package main

// var useInstance = []user{}
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
}

//
//func (u *user) getUser(s map[string]any) user {
//	err := json.NewDecoder().Decode(s)
//
//}

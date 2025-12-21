package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
}

//var users []User

//func (u *User) GetUsers(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	err := json.NewEncoder(w).Encode(users)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	} else {
//		w.WriteHeader(http.StatusOK)
//	}
//	return
//}

//func isValidUser(u User) bool {
//	if u.FirstName == "" || u.LastName == "" || u.Email == "" {
//		return false
//	} else {
//		return true
//	}
//	//if u.FirstName == "" {
//	//	return "Firstname is a required field!!"
//	//} else if u.LastName == "" {
//	//	return "Lastname is a required field!!"
//	//} else if u.Email == "" {
//	//	return "Email is a required field!!"
//	//}
//}

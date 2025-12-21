package main

import (
	"BackendDevelopment/tutorial/models"
	"BackendDevelopment/tutorial/validators"
	"encoding/json"
	"net/http"
)

type app struct {
	port  string
	debug bool
}

// create app methods to perform specific requests.
var users []models.User

func (app *app) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

func (app *app) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload models.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		b, s := validators.IsValidUser(payload)
		if !b {
			http.Error(w, s, http.StatusBadRequest)
			return
		}
	}
	users = append(users, payload)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payload)
}

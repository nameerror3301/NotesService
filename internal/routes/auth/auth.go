package auth

import (
	"NotesService/internal/models"
	"NotesService/internal/routes"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// From decode to user data
func beforeCreate(r *http.Request) (string, string) {
	var u models.UserData
	json.NewDecoder(r.Body).Decode(&u)

	return u.Email, u.Pass
}

// STATUS: WORK (TESTED)
func SignUp(w http.ResponseWriter, r *http.Request) {
	email, pass := beforeCreate(r)
	if email == "" || pass == "" {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "Incorrect data - check email or password"))
		return
	}

	status, err := models.CreateUser(email, pass)
	if err != nil {
		logrus.Warnf("Error in new user registration logic - %s", err)
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusInternalServerError, "Internal error"))
		return
	}

	if !status {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusUnauthorized, "A user with this email already exists"))
		return
	}

	logrus.Infof("A new user was registered - %t", status)
	json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "Success registration"))
}

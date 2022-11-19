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
	status, err := models.CreateUser(beforeCreate(r))
	if err != nil {
		logrus.Warnf("Error in new user registration logic - %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(routes.RespStatus(1.0, http.StatusInternalServerError, "Internal error"))
		return
	}

	if !status {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(routes.RespStatus(1.0, http.StatusInternalServerError, "A user with this email already exists"))
		return
	}

	logrus.Infof("A new user was registered - %t", status)
	json.NewEncoder(w).Encode(routes.RespStatus(1.0, http.StatusOK, "Success registration"))
}

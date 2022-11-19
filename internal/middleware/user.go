package middleware

import (
	"NotesService/internal/models"
	"NotesService/internal/routes"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// STATUS: WORK (Tested)
func UserSetContentType(next http.HandlerFunc) http.HandlerFunc {
	/*
		Set content type from responce
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// STATUS: WORK (Tested)
func UserRequestLog(next http.HandlerFunc) http.HandlerFunc {
	/*
		Middleware from logging user request
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("Path - [%s] - Method - [%s] Body - %v", r.URL.Path, r.Method, r.Body)
		next(w, r)
	}
}

// STATUS: WORK (Tested)
func UserMethodCheck(next http.HandlerFunc, method ...string) http.HandlerFunc {
	/*
		Middleware from check user method
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		if !isMethod(r, method...) {
			logrus.Warnf("The user uses the wrong method for this endpoint - [%s]", r.Method)

			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(routes.RespStatus(1.0, http.StatusMethodNotAllowed, "Incorrect Method"))
			return
		}
		next(w, r)
	}
}

// STATUS: WORK (Tested)
func isMethod(r *http.Request, method ...string) bool {
	var status []bool
	for _, val := range method {
		if r.Method == val {
			status = append(status, true)
		} else {
			status = append(status, false)
		}
	}

	for _, val := range status {
		if val {
			return true
		}
	}

	return false
}

// STATUS: WORK (Tested)
func UserCheckContent(next http.HandlerFunc) http.HandlerFunc {
	/*
		Middleware from check user content-type
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			logrus.Warnf("User sent an invalid data type - [%s]", r.Header.Get("Content-Type"))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(routes.RespStatus(1.0, http.StatusBadRequest, "Incorrect Content-Type"))
			return
		}
		next(w, r)
	}
}

func UserBasicAuth(next http.HandlerFunc) http.HandlerFunc {
	/*
		Middleware from auth user
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		email, pass, ok := r.BasicAuth()
		if ok && !models.IsUserData(email, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(routes.RespStatus(1.0, http.StatusUnauthorized, "Bad auth"))
			return
		} else {
			next(w, r)
		}
	}
}

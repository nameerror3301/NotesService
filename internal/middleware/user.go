package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// STATUS: WORK (Tested)
func UserRequestLog(next http.HandlerFunc, path string) http.HandlerFunc {
	/*
		Middleware from logging user request
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		if r.URL.Path != path {
			logrus.Warnf("User requested a page that does not exist - [%s]", r.URL.Path)
			http.Error(w, "Sorry, the page you requested does not exist", http.StatusNotFound)
			return
		}

		logrus.Infof("Path - [%s] - Method - [%s] Body - %v", r.URL.Path, r.Method, r.Body)
		next(w, r)
	}
}

// STATUS: NOTWORK
func UserMethodCheck(next http.HandlerFunc, method ...string) http.HandlerFunc {
	/*
		Middleware from check user method
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		// Add logic to work with multiple methods
		if r.Method != method {
			logrus.Warnf("The user uses the wrong method for this endpoint - [%s]", r.Method)
			http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

// STATUS: WORK (Tested)
func UserCheckContent(next http.HandlerFunc) http.HandlerFunc {
	/*
		Middleware from check user content-type
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		if r.Header.Get("Content-Type") != "application/json" {
			logrus.Warnf("User sent an invalid data type - [%s]", r.Header.Get("Content-Type"))
			http.Error(w, "Incorrect Content-Type", http.StatusBadRequest)
			return
		}
		next(w, r)
	}
}

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

// STATUS: WORK (tested)
func UserMethodCheck(next http.HandlerFunc, method ...string) http.HandlerFunc {
	/*
		Middleware from check user method
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if !isMethod(r, method...) {
			logrus.Warnf("The user uses the wrong method for this endpoint - [%s]", r.Method)
			http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

// STATUS: WORK (tested)
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
		w.Header().Set("Content-Type", "text/plain")

		if r.Header.Get("Content-Type") != "application/json" {
			logrus.Warnf("User sent an invalid data type - [%s]", r.Header.Get("Content-Type"))
			http.Error(w, "Incorrect Content-Type", http.StatusBadRequest)
			return
		}
		next(w, r)
	}
}

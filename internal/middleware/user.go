package middleware

import (
	"NotesService/internal/models"
	"NotesService/internal/routes"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Логирование запросов от пользователя
func UserRequestLog(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("Path - [%s] - Method - [%s]", r.URL.Path, r.Method)
		next(w, r)
	}
}

// Проверка метода запроса
func UserMethodCheck(next http.HandlerFunc, method ...string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if !isMethod(r, method...) {
			logrus.Warnf("The user uses the wrong method for this endpoint - [%s]", r.Method)

			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusMethodNotAllowed, "Incorrect Method"))
			return
		}
		next(w, r)
	}
}

// Вспомогательная функция для проверки метода запроса
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

// Проверка контента который отправляет пользователь
func UserCheckContent(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			logrus.Warnf("User sent an invalid data type - [%s]", r.Header.Get("Content-Type"))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "Incorrect Content-Type"))
			return
		}
		next(w, r)
	}
}

// Базовая аутентификация пользователя средствами BasicAuth
func UserBasicAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		email, pass, ok := r.BasicAuth()
		if ok && models.IsUserData(email, pass) {
			next(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusUnauthorized, "Bad auth"))
			return
		}
	}
}

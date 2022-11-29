package routes

import (
	"encoding/json"
	"net/http"
)

// Для формирования запросов
type resp struct {
	Api         float64     `json:"api_version"`
	Status      int         `json:"status"`
	Description interface{} `json:"description"`
}

// Функция для формирования ответа пользователю
func RespStatus(w http.ResponseWriter, api float64, status int, description interface{}) []resp {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var res []resp
	res = append(res, resp{
		Api:         api,
		Status:      status,
		Description: description,
	})
	return res
}

// Стартовая страница
func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "This is home page!"))
}

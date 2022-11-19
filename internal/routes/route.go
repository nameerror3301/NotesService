package routes

import (
	"encoding/json"
	"net/http"
)

type resp struct {
	Api         float32 `json:"api_version"`
	Status      int     `json:"status"`
	Description string  `json:"description"`
}

func RespStatus(api float32, status int, description string) []resp {
	var res []resp
	res = append(res, resp{
		Api:         api,
		Status:      status,
		Description: description,
	})
	return res
}

// STATUS: WORK
func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RespStatus(1.0, http.StatusOK, "This is home page!"))
}

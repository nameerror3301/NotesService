package routes

import (
	"net/http"
)

const (
	StatusOK   = http.StatusOK
	BadRequest = http.StatusBadRequest
	NotFound   = http.StatusNotFound
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Start Page!"))
}

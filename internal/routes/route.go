package routes

import (
	"net/http"
)

// STATUS: WORK
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Start Page!"))
}

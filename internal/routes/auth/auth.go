package auth

import "net/http"

// STATUS: Not implemented
func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignIn"))
}

// STATUS: Not implemented
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignUp"))
}

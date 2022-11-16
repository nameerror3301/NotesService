package run

import (
	"net/http"
	"time"

	middle "NotesService/internal/middleware"
	routes "NotesService/internal/routes"
	auth "NotesService/internal/routes/auth"
)

func Run() error {
	// Server config
	serv := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// http.HandleFunc("/api") -> There will be Swagger documentation
	http.HandleFunc("/api/v1", middle.UserMethodCheck(middle.UserRequestLog(routes.HomePage, "/api/v1"), http.MethodGet))
	http.HandleFunc("/api/v1/signIn", middle.UserMethodCheck(middle.UserCheckContent(middle.UserRequestLog(auth.SignIn, "/api/v1/signIn")), http.MethodPost))
	http.HandleFunc("/api/v1/signUp", middle.UserMethodCheck(middle.UserCheckContent(middle.UserRequestLog(auth.SignUp, "/api/v1/signUp")), http.MethodPost))

	if err := serv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

package run

import (
	"net/http"
	"time"

	middle "NotesService/internal/middleware"
	routes "NotesService/internal/routes"
	auth "NotesService/internal/routes/auth"
	"NotesService/internal/routes/notes"
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
	http.HandleFunc("/api/v1", middle.UserMethodCheck(middle.UserRequestLog(routes.HomePage), http.MethodGet, http.MethodPost))
	http.HandleFunc("/api/v1/signUp", middle.UserMethodCheck(middle.UserCheckContent(auth.SignUp), http.MethodPost))
	http.HandleFunc("/api/v1/notes", middle.UserMethodCheck(middle.UserBasicAuth(middle.UserRequestLog(notes.FindAllNotes)), http.MethodGet))
	http.HandleFunc("/api/v1/notes/create", middle.UserMethodCheck(middle.UserBasicAuth(middle.UserRequestLog(notes.CreateNote)), http.MethodPost))
	http.HandleFunc("/api/v1/notes/upload", middle.UserMethodCheck(middle.UserBasicAuth(middle.UserRequestLog(notes.UploadNote)), http.MethodPut))

	if err := serv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

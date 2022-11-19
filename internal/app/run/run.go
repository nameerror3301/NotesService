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
	http.HandleFunc("/api/v1", middle.UserSetContentType(middle.UserMethodCheck(middle.UserRequestLog(routes.HomePage), http.MethodGet, http.MethodPost)))
	http.HandleFunc("/api/v1/signUp", middle.UserSetContentType(middle.UserMethodCheck(middle.UserCheckContent(auth.SignUp), http.MethodPost)))
	// http.HandleFunc("/api/v1/notes")        // Получение всех заметок
	// http.HandleFunc("/api/v1/notes/create") // Создание заметок
	// http.HandleFunc("/api/v1/notes/upload") // Удаление и обновление заметок

	if err := serv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

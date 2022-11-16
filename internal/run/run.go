package run

import (
	"net/http"
	"time"

	routes "NotesService/internal/routes"
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

	http.HandleFunc("/", routes.HomePage)

	if err := serv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

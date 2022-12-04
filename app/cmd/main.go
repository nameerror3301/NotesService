package main

import (
	run "NotesService/internal/app/run"

	logrus "github.com/sirupsen/logrus"
)

func init() {
	// Минимальная конфигурация формата логов
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
}

func main() {
	logrus.Info("Server up successful!")
	if err := run.Run(); err != nil {
		logrus.Fatalf("Err server up - %s", err)
	}
}

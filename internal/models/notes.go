package models

import (
	"math/rand"

	"github.com/sirupsen/logrus"
)

type NotesData struct {
	Email string `json:"-"`
	Id    int    `json:"id"`
	Name  string `json:"notes_name"`
	Value string `json:"notes_content"`
}

var notes []NotesData

/*
	Getting all notes from a particular user (The list of notes is determined by email)
	Email is not visible to the user when all notes are received
*/

func FindAll(email string) []NotesData {

	var data []NotesData

	for _, val := range notes {
		if val.Email == email {
			data = append(data, NotesData{
				Email: val.Email,
				Id:    val.Id,
				Name:  val.Value,
				Value: val.Value,
			})
		}
	}

	// Проверка на наличие у пользователя заметок
	if data == nil {
		return nil
	}

	return data
}

// Creating a note
func CreateNote(email string, name string, value string) {
	notes = append(notes, NotesData{
		Email: email,
		Id:    rand.Intn(10000),
		Name:  name,
		Value: value,
	})

	logrus.Infof("User - %s --> Create Notes", email)
}

// Upload data in notes

// Delite notes

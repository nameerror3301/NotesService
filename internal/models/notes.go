package models

import (
	"time"

	"github.com/sirupsen/logrus"
)

type NotesData struct {
	Email string `json:"-"`
	Id    int    `json:"id"`
	Name  string `json:"notes_name"`
	Value string `json:"notes_content"`
	// TTL
	Date time.Time `json:"last_update"`
}

var notes []NotesData

func numbering(data []NotesData) int {
	if data == nil {
		return 1
	}

	num := data[len(data)-1]

	return num.Id + 1
}

/*
	Getting all notes from a particular user (The list of notes is determined by email)
	Email is not visible to the user when all notes are received
*/
// Find all WORK: TESTED
func FindAll(email string) []NotesData {
	var data []NotesData

	for _, val := range notes {
		if val.Email == email {
			data = append(data, NotesData{
				Email: val.Email,
				Id:    val.Id,
				Name:  val.Name,
				Value: val.Value,
				Date:  val.Date,
			})
		}
	}

	// Checking for user notes
	if data == nil {
		return nil
	}

	return data
}

// Find by id WORK: TESTED
func FindById(email string, id int) []NotesData {
	var data []NotesData
	for _, val := range notes {
		if val.Email == email {
			if val.Id == id {
				data = append(data, NotesData{
					Email: val.Email,
					Id:    val.Id,
					Name:  val.Name,
					Value: val.Value,
					Date:  val.Date,
				})
				return data
			}
		}
	}
	return nil
}

// Creating a note WORK: TESTED
func CreateNote(email string, name string, value string) {
	notes = append(notes, NotesData{
		Email: email,
		Id:    numbering(notes),
		Name:  name,
		Value: value,
		Date:  time.Now(),
	})
	logrus.Infof("%s --> Create Notes", email)
}

// Upload notes WORK: TESTED
func UploadNote(email string, id int, newname string, newvalue string) bool {
	for idx, val := range notes {
		if val.Email == email {
			if val.Id == id {
				notes[idx].Name = newname
				notes[idx].Value = newvalue
				return true
			}
		}
	}
	return false
}

// Delete notes WORK: TESTED
func DeliteNote(email string, id int) bool {
	status := false

	for idx, val := range notes {
		if val.Email == email {
			if val.Id == id {
				notes = append(notes[:idx], notes[idx+1:]...)
				status = true
			}
		}
	}

	for idx := range notes {
		notes[idx].Id = idx + 1
	}
	return status
}

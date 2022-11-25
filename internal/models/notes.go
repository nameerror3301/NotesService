package models

import (
	"github.com/sirupsen/logrus"
)

type NotesData struct {
	Email string `json:"-"`
	Id    int    `json:"id"`
	Name  string `json:"notes_name"`
	Value string `json:"notes_content"`
	// TTL
	// Date create
}

var notes []NotesData

func Numbering(data []NotesData) int {
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

func FindAll(email string) []NotesData {
	var data []NotesData

	for _, val := range notes {
		if val.Email == email {
			data = append(data, NotesData{
				Email: val.Email,
				Id:    val.Id,
				Name:  val.Name,
				Value: val.Value,
			})
		}
	}

	// Checking for user notes
	if data == nil {
		return nil
	}

	return data
}

// Creating a note
func CreateNote(email string, name string, value string) {
	num := Numbering(notes)

	notes = append(notes, NotesData{
		Email: email,
		Id:    num,
		Name:  name,
		Value: value,
	})
	logrus.Infof("%s --> Create Notes", email)
}

// Upload data in notes (PUT)
func UploadNote(email string, id int, newname string, newvalue string) {

}

// Delite notes (DELITE)
func DeliteNote(email string, id int) {
	/*
		Удаление будет реализованно средствами среза слайса структур
	*/
}

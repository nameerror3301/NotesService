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

// Общая структура для хранения заметок
var notes []NotesData

func numbering(data []NotesData) int {
	if data == nil {
		return 1
	}

	num := data[len(data)-1]

	return num.Id + 1
}

/*
	При создании заметок пользователем,
		запускать независимую от основого потока выполнения функцию таймер которая через
			определенное время будет проверять заметки с истекшим TTL и удалять их. Благо,
				функция для удаления уже реализованна!
*/

// Сортировка по ID всех заметок (Пузырьковая)
func sortById(data []NotesData, sortQuery string) []NotesData {
	for i := 0; i < len(data)-1; i++ {
		for j := i; j < len(data); j++ {
			if sortQuery == "ASC" {
				if data[i].Id > data[j].Id {
					data[i], data[j] = data[j], data[i]
				}
			} else {
				if data[i].Id < data[j].Id {
					data[i], data[j] = data[j], data[i]
				}
			}
		}
	}
	return data
}

// Получение всех заметок (Не сортированных)
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

	return data
}

// Получение всех заметок (Отсортированных)
func FindAllSort(email string, querySort string) []NotesData {
	data := FindAll(email)

	return sortById(data, querySort)
}

// Получение по ID
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

// Создание заметок
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

// Обновление или дополнение заметок
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

// Удаление заметок по их ID
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

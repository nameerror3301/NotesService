package models

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type NotesData struct {
	Email   string    `json:"-"`
	Id      int       `json:"id"`
	Name    string    `json:"notes_name"`
	Value   string    `json:"notes_content"`
	TTL     string    `json:"ttl_date_to_delete"`
	RemTime time.Time `json:"-"`
	Date    time.Time `json:"upload_date"`
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
				TTL:   val.TTL,
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

	// INFO: Реализовать бинарный поиск по слайсу структур
	var data []NotesData
	for _, val := range notes {
		if val.Email == email {
			if val.Id == id {
				data = append(data, NotesData{
					Email: val.Email,
					Id:    val.Id,
					Name:  val.Name,
					Value: val.Value,
					TTL:   val.TTL,
					Date:  val.Date,
				})
				return data
			}
		}
	}
	return nil
}

// Создание заметок
func CreateNote(email string, name string, value string, ttl string) bool {
	parceTTL, err := time.Parse("2006/01/02  15:04:05", ttl)
	if err != nil {
		logrus.Warnf("The user transmitted an incorrect date - %s", err)
		return false
	}

	if parceTTL.UTC().Before(time.Now()) || parceTTL.UTC().Equal(time.Now()) {
		logrus.Warn("The user is trying to pass the date of the past time")
		return false
	} else {
		notes = append(notes, NotesData{
			Email:   email,
			Id:      numbering(notes),
			Name:    name,
			Value:   value,
			TTL:     ttl,
			RemTime: parceTTL,
			Date:    time.Now(),
		})

		logrus.Infof("%s --> Create Notes", email)
		return true
	}
}

// Обновление или дополнение заметок
func UploadNote(email string, id int, newname string, newvalue string) bool {
	for idx, val := range notes {
		if val.Email == email {
			if val.Id == id {
				notes[idx].Name = newname
				notes[idx].Value = newvalue
				notes[idx].Date = time.Now()
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

// Пофиксить (Не работает)
func CheckingExpirationDates() {
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go func() {
			timer := time.NewTimer(5 * time.Second)
			<-timer.C
			for _, val := range notes {
				if val.RemTime.UTC().Equal(time.Now()) || val.RemTime.UTC().After(time.Now()) {
					if DeliteNote(val.Email, val.Id) {
						logrus.Info("The note expired and was deleted - %d", val.Id)
					}
				}
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

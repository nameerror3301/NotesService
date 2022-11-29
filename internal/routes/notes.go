package routes

import (
	"NotesService/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Получение всех заметок без сортировки, получение с сортировкой, получение заметки по ID.
func FindAllNotesOrById(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()

	idStr := r.URL.Query().Get("id")
	querySort := r.URL.Query().Get("sort")

	if querySort != "" {
		if querySort == "ASC" || querySort == "DESC" {
			if data := models.FindAllSort(email, querySort); data == nil {
				json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "The note with the specified id was not found"))
				return
			} else {
				json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, &data))
				return
			}

		}
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusBadRequest, "Incorrect parameters for sorting"))
		return
	}

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil && id > 0 {
			logrus.Warnf("The user sends an invalid value - %s", idStr)
			json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusBadRequest, "You sent the wrong parameter"))
			return
		}

		if data := models.FindById(email, id); data == nil {
			json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "The note with the specified id was not found"))
			return
		} else {
			json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, &data))
			return
		}
	}

	if data := models.FindAll(email); data == nil {
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "You don't have notes!"))
		return
	} else {
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, &data))
		return
	}
}

// Создание заметки
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()

	json.NewDecoder(r.Body).Decode(&note)

	if note.Name == "" {
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusBadRequest, "Check the data entered correctly, fields should not be empty when creating the note!"))
		return
	}

	models.CreateNote(email, note.Name, note.Value)
	json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "Success create!"))
}

// Обновление заметки
func UploadNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()
	json.NewDecoder(r.Body).Decode(&note)

	// Check nil in fields
	if note.Name == "" || note.Value == "" {
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusBadRequest, "Invalid value in the fields, check fields!"))
		return
	}

	if status := models.UploadNote(email, note.Id, note.Name, note.Value); status {
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "Upload success!"))
		return
	} else {
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "No notes with this id were found!"))
		return
	}

}

// Удаление заметки
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil && id > 0 {
		logrus.Warnf("The user sends an invalid value - %s", idStr)
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusBadRequest, "You sent the wrong parameter"))
		return
	}

	if status := models.DeliteNote(email, id); status {
		json.NewEncoder(w).Encode(RespStatus(w, 1.0, http.StatusOK, "Delete success!"))
		return
	}
}

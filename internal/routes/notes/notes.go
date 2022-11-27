package notes

import (
	"NotesService/internal/models"
	"NotesService/internal/routes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Объединяет в себе две функции получение по id и получение всех заметок
func FindAllNotesOrById(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()

	idStr := r.URL.Query().Get("id")

	fmt.Println(idStr)

	if idStr == "" {
		if data := models.FindAll(email); data == nil {
			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "You don't have notes!"))
			return
		} else {
			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, &data))
			return
		}
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil && id > 0 {
			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "You sent the wrong parameter"))
			return
		}

		if data := models.FindById(email, id); data == nil {
			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "The note with the specified id was not found"))
			return
		} else {
			json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, &data))
			return
		}
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()

	json.NewDecoder(r.Body).Decode(&note)

	// Check nil in name notes
	if note.Name == "" {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "Check the data entered correctly, fields should not be empty when creating the note!"))
	}

	models.CreateNote(email, note.Name, note.Value)
	json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "Success create!"))
}

// Заготовка под обновление заметки
func UploadNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()
	json.NewDecoder(r.Body).Decode(&note)

	// Check nil in fields
	if note.Name == "" || note.Value == "" {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "Invalid value in the fields, check fields!"))
		return
	}

	/*
		The function will return false if the user has no notes at all
			or no notes with the specified id are found
	*/
	if status := models.UploadNote(email, note.Id, note.Name, note.Value); status {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "Upload success!"))
		return
	} else {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "No notes with this id were found!"))
		return
	}

}

// Заготовка под удаление заметки
func DeliteNote(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil && id > 0 {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "You sent the wrong parameter"))
	}

	if status := models.DeliteNote(email, id); status {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "Delete success!"))
		return
	}
}

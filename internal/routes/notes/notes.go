package notes

import (
	"NotesService/internal/models"
	"NotesService/internal/routes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Receiving notes, sorting, retrieving by ID
func FindAllNotesOrById(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()

	idStr := r.URL.Query().Get("id")
	querySort := r.URL.Query().Get("sort")

	if querySort != "" {
		if querySort == "ASC" || querySort == "DESC" {
			if data := models.FindAllSort(email, querySort); data == nil {
				json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "The note with the specified id was not found"))
			} else {
				json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, &data))
				return

			}

		}
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "Incorrect parameters for sorting"))
		return
	}

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil && id > 0 {
			logrus.Warnf("The user sends an invalid value - %s", idStr)
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

	if data := models.FindAll(email); data == nil {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "You don't have notes!"))
		return
	} else {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, &data))
		return
	}
}

// To create notes
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()

	json.NewDecoder(r.Body).Decode(&note)

	// Check nil in name notes
	if note.Name == "" {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "Check the data entered correctly, fields should not be empty when creating the note!"))
		return
	}

	models.CreateNote(email, note.Name, note.Value)
	json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "Success create!"))
}

// To update notes by ID
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

// To delete notes by ID
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil && id > 0 {
		logrus.Warnf("The user sends an invalid value - %s", idStr)
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "You sent the wrong parameter"))
		return
	}

	if status := models.DeliteNote(email, id); status {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "Delete success!"))
		return
	}
}

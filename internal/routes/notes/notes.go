package notes

import (
	"NotesService/internal/models"
	"NotesService/internal/routes"
	"encoding/json"
	"net/http"
)

func FindAllNotes(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()
	if data := models.FindAll(email); data == nil {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "You don't have notes"))
	} else {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, &data))
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()

	json.NewDecoder(r.Body).Decode(&note)

	if note.Name == "" {
		json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusBadRequest, "Check the data entered correctly, fields should not be empty when creating the note"))
	}

	models.CreateNote(email, note.Name, note.Value)
	json.NewEncoder(w).Encode(routes.RespStatus(w, 1.0, http.StatusOK, "Success create"))

}

// Заготовка под обновление заметки
func UploadNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()

	json.NewDecoder(r.Body).Decode(&note)

	models.UploadNote(email, note.Id, note.Name, note.Value)
}

// Заготовка под удаление заметки
func DeliteNote(w http.ResponseWriter, r *http.Request) {
	var note models.NotesData
	email, _, _ := r.BasicAuth()

	json.NewDecoder(r.Body).Decode(&note)

	models.DeliteNote(email, note.Id)
}

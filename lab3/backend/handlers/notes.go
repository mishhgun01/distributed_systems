package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"notes/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetNotes возвращает все заметки
func GetNotes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var notes []models.Note
		db.Find(&notes)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	}
}

// CreateNote создает новую заметку
func CreateNote(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note models.Note
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.Create(&note)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
	}
}

// UpdateNote обновляет существующую заметку
func UpdateNote(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note models.Note
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Проверяем, существует ли заметка
		var existingNote models.Note
		if err := db.First(&existingNote, note.ID).Error; err != nil {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}

		// Обновляем заметку
		existingNote.Title = note.Title
		existingNote.Content = note.Content
		db.Save(&existingNote)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(existingNote)
	}
}

// DeleteNote deletes a note from the database
func DeleteNote(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var note models.Note
		if err := db.First(&note, id).Error; err != nil {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}

		db.Delete(&note)
		w.WriteHeader(http.StatusNoContent)
	}
}

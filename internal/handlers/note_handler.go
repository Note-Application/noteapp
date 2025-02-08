package handlers

import (
	"encoding/json"
	"net/http"
	"noteapp/internal/models"
	"noteapp/internal/services"
	"github.com/gorilla/mux"
)

// CreateNote handles creating a new note
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.CreateNote(note)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// GetNoteByID retrieves a note by its ID
func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	note, err := services.GetNoteByID(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(note)
}

// GetNotesByUserID retrieves all notes for a specific user
func GetNotesByUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]
	notes, err := services.GetNotesByUserID(userID)
	if err != nil {
		http.Error(w, "No notes found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(notes)
}

// UpdateNote updates an existing note
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.UpdateNote(id, note)
	if err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(note)
}

// DeleteNote deletes a note by its ID
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := services.DeleteNote(id)
	if err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
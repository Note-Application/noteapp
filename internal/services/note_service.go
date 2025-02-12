package services

import (
	"noteapp/internal/models"
	"noteapp/internal/repositories"
	"noteapp/pkg/db"
	"log"
)

// CreateNote creates a new note in the database and returns the ID of the newly created note
func CreateNote(note models.Note) (int, error) {
    id, err := repositories.CreateNote(db.DB, note)
    if err != nil {
        log.Println("Error creating note:", err)
        return 0, err
    }
    return id, nil
}

// GetNoteByID retrieves a note by its ID
func GetNoteByID(id string) (models.Note, error) {
	note, err := repositories.GetNoteByID(db.DB, id)
	if err != nil {
		log.Println("Error getting note by ID:", err)
		return note, err
	}
	return note, nil
}

// GetNotesByUserID retrieves all notes for a specific user
func GetNotesByUserID(userID string) ([]models.Note, error) {
	notes, err := repositories.GetNotesByUserID(db.DB, userID)
	if err != nil {
		log.Println("Error getting notes by user ID:", err)
		return nil, err
	}
	return notes, nil
}

// UpdateNote updates an existing note
func UpdateNote(id string, note models.Note) error {
	err := repositories.UpdateNote(db.DB, id, note)
	if err != nil {
		log.Println("Error updating note:", err)
		return err
	}
	return nil
}

// DeleteNote deletes a note by ID
func DeleteNote(id string) error {
	err := repositories.DeleteNote(db.DB, id)
	if err != nil {
		log.Println("Error deleting note:", err)
		return err
	}
	return nil
}
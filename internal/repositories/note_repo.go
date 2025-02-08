package repositories

import (
	"database/sql"
	"noteapp/internal/models"
	"log"
)

// CreateNote creates a new note in the database
func CreateNote(db *sql.DB, note models.Note) error {
	_, err := db.Exec("INSERT INTO notes (user_id, title, content) VALUES ($1, $2, $3)", note.UserID, note.Title, note.Content)
	if err != nil {
		log.Println("Error creating note:", err)
		return err
	}
	return nil
}

// GetNoteByID retrieves a note by its ID
func GetNoteByID(db *sql.DB, id string) (models.Note, error) {
	var note models.Note
	err := db.QueryRow("SELECT id, user_id, title, content, created_at, updated_at FROM notes WHERE id = $1", id).Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return note, err
	}
	return note, nil
}

// GetNotesByUserID retrieves all notes for a specific user
func GetNotesByUserID(db *sql.DB, userID string) ([]models.Note, error) {
	rows, err := db.Query("SELECT id, user_id, title, content, created_at, updated_at FROM notes WHERE user_id = $1", userID)
	if err != nil {
		log.Println("Error retrieving notes by user ID:", err)
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			log.Println("Error scanning note:", err)
			continue
		}
		notes = append(notes, note)
	}
	return notes, nil
}

// UpdateNote updates an existing note
func UpdateNote(db *sql.DB, id string, note models.Note) error {
	_, err := db.Exec("UPDATE notes SET title=$1, content=$2, updated_at=NOW() WHERE id=$3", note.Title, note.Content, id)
	if err != nil {
		log.Println("Error updating note:", err)
		return err
	}
	return nil
}

// DeleteNote deletes a note by ID
func DeleteNote(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM notes WHERE id=$1", id)
	if err != nil {
		log.Println("Error deleting note:", err)
		return err
	}
	return nil
}
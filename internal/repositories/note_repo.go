package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"noteapp/internal/models"
	"noteapp/redis"
)

// CreateNote creates a new note in the database and returns the ID of the newly created note
func CreateNote(db *sql.DB, note models.Note) (int, error) {
	var id int
	query := `INSERT INTO notes (user_id, title, content) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, note.UserID, note.Title, note.Content).Scan(&id)
	if err != nil {
		log.Println("Error creating note:", err)
		return 0, err
	}
	return id, nil
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
	rows, err := db.Query("SELECT id, user_id, title, content, created_at, updated_at FROM notes WHERE user_id = $1 ORDER BY created_at ASC", userID)
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

		// Check if updated note exists in Redis
		redisData, err := redis.GetNote(fmt.Sprintf("%d", note.ID))
		if err == nil && redisData != "" { // Redis data exists
			var updatedNote map[string]string
			if err := json.Unmarshal([]byte(redisData), &updatedNote); err == nil {
				// Override title and content from Redis
				note.Title = updatedNote["title"]
				note.Content = updatedNote["content"]
			} else {
				log.Println("Failed to parse JSON from Redis for note ID:", note.ID)
			}
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

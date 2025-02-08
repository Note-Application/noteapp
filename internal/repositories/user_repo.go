package repositories

import (
	"database/sql"
	"noteapp/internal/models"
	"log"
)

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users (email, name, profile_pic) VALUES ($1, $2, $3)", user.Email, user.Name, user.ProfilePic)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user by email from the database
func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, email, name, profile_pic, created_at FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Name, &user.ProfilePic, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(db *sql.DB, email string, user models.User) error {
	_, err := db.Exec("UPDATE users SET name=$1, profile_pic=$2 WHERE email=$3", user.Name, user.ProfilePic, email)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	return nil
}

// DeleteUser deletes a user by email
func DeleteUser(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM users WHERE email=$1", email)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}
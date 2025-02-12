package services

import (
	"noteapp/internal/models"
	"noteapp/internal/repositories"
	"noteapp/pkg/db"
	"log"
)

// CreateUser creates a new user in the database and returns the ID of the newly created user
func CreateUser(user models.User) (int, error) {
    id, err := repositories.CreateUser(db.DB, user)
    if err != nil {
        log.Println("Error creating user:", err)
        return 0, err
    }
    return id, nil
}

// GetUserByEmail retrieves a user by email from the database
func GetUserByEmail(email string) (models.User, error) {
	user, err := repositories.GetUserByEmail(db.DB, email)
	if err != nil {
		log.Println("Error getting user by email:", err)
		return user, err
	}
	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(email string, user models.User) error {
	err := repositories.UpdateUser(db.DB, email, user)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	return nil
}

// DeleteUser deletes a user by email
func DeleteUser(email string) error {
	err := repositories.DeleteUser(db.DB, email)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}
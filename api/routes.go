package api

import (
	"noteapp/internal/handlers"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(r *mux.Router) {
	// User routes
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{email}", handlers.GetUserByEmail).Methods("GET")
	r.HandleFunc("/users/{email}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{email}", handlers.DeleteUser).Methods("DELETE")

	// Note routes
	r.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id}", handlers.GetNoteByID).Methods("GET")
	r.HandleFunc("/notes/user/{user_id}", handlers.GetNotesByUserID).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.UpdateNote).Methods("PUT")
	r.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")
}

package main

import (
	"log"
	"net/http"
	"noteapp/api"
	"noteapp/config"
	"noteapp/pkg/db"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	//config.LoadConfig()

	// Initialize database connection
	db.InitDB()

	// Setup routes
	r := mux.NewRouter()
	api.RegisterRoutes(r)

	// Apply CORS middleware
	corsHandler := handlers.CORS(
		// Allow all origins
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Accept"}),
	)(r)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server starting on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the PostgreSQL database connection
func InitDB() {
	// Get database details from environment variables
	dbURL := os.Getenv("DATABASE_URL") // For Aiven PostgreSQL URI
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	var connStr string
	if dbURL != "" {
		// Use Aiven connection URL if provided
		connStr = dbURL
	} else {
		// Use manual connection string if Aiven URL is not set
		connStr = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbUser, dbPassword, dbHost, dbPort, dbName,
		)
	}

	connStr = "postgres://avnadmin:AVNS_wrkNseB_iSwng4azMWD@noteappdb-zoinme.k.aivencloud.com:22802/noteapp_db?sslmode=require"
	// Open database connection
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Ping the database to check connectivity
	if err := DB.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	log.Println("✅ Successfully connected to the database!")
}

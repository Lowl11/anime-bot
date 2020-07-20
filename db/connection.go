package db

import (
	"database/sql"
	"fmt"
	"log"

	// Library for using PostgreSQL
	_ "github.com/lib/pq"
)

// Database connected to PostgreSQL
var Database *sql.DB

// InitializeDatabase - connect object to database
func InitializeDatabase(username, password, dbName string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbName)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Database successfully connected")
	Database = database
}

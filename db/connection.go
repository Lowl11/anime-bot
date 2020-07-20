package db

import (
	"database/sql"
	"log"

	// Library for using PostgreSQL
	_ "github.com/lib/pq"
)

// Database connected to PostgreSQL
var Database *sql.DB

// InitializeDatabase - connect object to database
func InitializeDatabase() {
	connectionString := "user=postgres password=qwerty11 dbname=anime-feed sslmode=disable"
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}
	Database = database
}

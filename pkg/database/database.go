package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"payment-gateway/config"
)

// NewPostgresDB initializes a connection to a PostgreSQL database
func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	// Connection string for PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL database!")
	return db, nil
}

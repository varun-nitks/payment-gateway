// database/database.go
package database

import (
	_ "cloud.google.com/go/cloudsqlconn" // Google Cloud SQL connection
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"payment-gateway/config"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	var connStr string

	if cfg.Environment == "gcp" && cfg.UseCloudSQLAuth {
		// Use Cloud SQL Auth Proxy for GCP
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBName)
		_, err := sql.Open("cloudsqlpostgres", cfg.Instance+":"+dsn)
		if err != nil {
			return nil, err
		}
	} else {
		// Local or non-Cloud SQL connection
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			"localhost", "5432", cfg.DBUser, cfg.DBPassword, cfg.DBName)
	}

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

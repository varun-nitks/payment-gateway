// config/config.go
package config

import (
	"os"
)

type Config struct {
	DBUser          string
	DBPassword      string
	DBName          string
	Host            string
	Port            string
	Instance        string
	Environment     string
	UseCloudSQLAuth bool // Use true if connecting to Cloud SQL on GCP
}

func LoadConfig() *Config {
	return &Config{
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		Host:            os.Getenv("DB_HOST"),
		Port:            os.Getenv("DB_PORT"),
		Instance:        os.Getenv("INSTANCE_CONNECTION_NAME"),
		Environment:     os.Getenv("ENVIRONMENT"), // e.g., "local" or "gcp"
		UseCloudSQLAuth: os.Getenv("USE_CLOUD_SQL_AUTH") == "true",
	}
}

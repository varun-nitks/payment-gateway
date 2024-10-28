// config/config.go

package config

import (
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	Instance   string
}

func LoadConfig() *Config {
	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Instance:   os.Getenv("INSTANCE_CONNECTION_NAME"),
	}
}

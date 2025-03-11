// internal/config/config.go
package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func Load() (*Config, error) {
	// Default values
	serverPort := 8080
	dbHost := "localhost"
	dbPort := 9000
	dbName := "default"
	dbUser := "default"
	dbPass := ""

	// Override with environment variables if provided
	if port, exists := os.LookupEnv("SERVER_PORT"); exists {
		if p, err := strconv.Atoi(port); err == nil {
			serverPort = p
		}
	}

	if host, exists := os.LookupEnv("DB_HOST"); exists {
		dbHost = host
	}

	if port, exists := os.LookupEnv("DB_PORT"); exists {
		if p, err := strconv.Atoi(port); err == nil {
			dbPort = p
		}
	}

	if name, exists := os.LookupEnv("DB_NAME"); exists {
		dbName = name
	}

	if user, exists := os.LookupEnv("DB_USER"); exists {
		dbUser = user
	}

	if pass, exists := os.LookupEnv("DB_PASS"); exists {
		dbPass = pass
	}

	return &Config{
		Server: ServerConfig{
			Port: serverPort,
		},
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			Database: dbName,
			Username: dbUser,
			Password: dbPass,
		},
	}, nil
}

package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	clickhouse "github.com/ClickHouse/clickhouse-go/v2"
)

var ClickHouseConn clickhouse.Conn

// InitClickHouse initializes the ClickHouse connection using environment variables or defaults.
func InitClickHouse() {
	// Read configuration from environment variables with defaults
	host := os.Getenv("CLICKHOUSE_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("CLICKHOUSE_PORT")
	if port == "" {
		port = "9000"
	}

	// Combine host and port into a single address
	address := fmt.Sprintf("%s:%s", host, port)

	database := os.Getenv("CLICKHOUSE_DATABASE")
	if database == "" {
		database = "default"
	}

	username := os.Getenv("CLICKHOUSE_USERNAME")
	if username == "" {
		username = "default"
	}

	password := os.Getenv("CLICKHOUSE_PASSWORD") // if not set, empty password is used

	var err error
	ClickHouseConn, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{address},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to clickhouse: %v", err)
	}

	// Ping to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ClickHouseConn.Ping(ctx); err != nil {
		log.Fatalf("clickhouse ping error: %v", err)
	}
	log.Println("Connected to ClickHouse successfully")
}

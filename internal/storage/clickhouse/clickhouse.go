// internal/storage/clickhouse/clickhouse.go
package clickhouse

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/shinshARK/clickhouse-api/internal/config"
	"github.com/shinshARK/clickhouse-api/internal/storage/repository"
)

func Connect(cfg config.DatabaseConfig) (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.Username,
			Password: cfg.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Debug: false,
	})

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	// Create tables if they don't exist
	if err := createTables(conn); err != nil {
		return nil, err
	}

	return conn, nil
}

func createTables(conn driver.Conn) error {
	// Create users table
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id String,
		username String,
		email String,
		created_at DateTime,
		updated_at DateTime
	) ENGINE = MergeTree()
	ORDER BY (id);
	`

	return conn.Exec(context.Background(), query)
}

func NewRepositories(db driver.Conn) *repository.Repositories {
	return &repository.Repositories{
		User: NewUserRepository(db),
	}
}

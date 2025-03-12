package repository

import (
	"context"
	"database/sql"

	"crud-fiber-clickhouse/config"
	"crud-fiber-clickhouse/models"
)

// CreateUser inserts a new user into ClickHouse.
func CreateUser(ctx context.Context, user models.User) error {
	// ClickHouse is append-only, so we assume INSERT operations.
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	return config.ClickHouseConn.Exec(ctx, query, user.ID, user.Name, user.Email)
}

// GetUser retrieves a user by ID.
func GetUser(ctx context.Context, id uint64) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ? LIMIT 1"
	var user models.User

	// QueryRow may be wrapped in a loop if you expect multiple rows
	if err := config.ClickHouseConn.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user's information.
func UpdateUser(ctx context.Context, user models.User) error {
	// ClickHouse is optimized for inserts and batch queries.
	// UPDATE operations are emulated by inserting a new version, but for this example we assume a simple update.
	query := "ALTER TABLE users UPDATE name = ?, email = ? WHERE id = ?"
	return config.ClickHouseConn.Exec(ctx, query, user.Name, user.Email, user.ID)
}

// DeleteUser deletes a user record.
func DeleteUser(ctx context.Context, id uint64) error {
	// In ClickHouse, DELETE operations are not typical.
	// We simulate deletion using an ALTER TABLE DELETE statement.
	query := "ALTER TABLE users DELETE WHERE id = ?"
	return config.ClickHouseConn.Exec(ctx, query, id)
}

// ListUsers retrieves all users (or a paginated list).
func ListUsers(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := config.ClickHouseConn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

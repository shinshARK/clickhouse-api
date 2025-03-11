// internal/storage/clickhouse/user.go
package clickhouse

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
	"github.com/shinshARK/clickhouse-api/internal/models"
)

type UserRepository struct {
	db driver.Conn
}

func NewUserRepository(db driver.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	// Generate a new UUID if not provided
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := `
	INSERT INTO users (id, username, email, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`

	return r.db.Exec(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `
	SELECT id, username, email, created_at, updated_at
	FROM users
	WHERE id = ?
	LIMIT 1
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("user not found")
	}

	var user models.User
	err = rows.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context, limit, offset int) ([]*models.User, error) {
	if limit <= 0 {
		limit = 10
	}

	query := `
	SELECT id, username, email, created_at, updated_at
	FROM users
	ORDER BY created_at DESC
	LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now()

	query := `
	ALTER TABLE users
	UPDATE username = ?, email = ?, updated_at = ?
	WHERE id = ?
	`

	return r.db.Exec(ctx, query,
		user.Username,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `
	ALTER TABLE users
	DELETE WHERE id = ?
	`

	return r.db.Exec(ctx, query, id)
}

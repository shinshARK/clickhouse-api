// internal/storage/repository.go
package repository

import (
	"context"

	"github.com/shinshARK/clickhouse-api/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type Repositories struct {
	User UserRepository
}

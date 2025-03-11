// internal/api/handlers/handlers.go
package handlers

import (
	"github.com/shinshARK/clickhouse-api/internal/storage/repository"
	"github.com/shinshARK/clickhouse-api/pkg/logger"
)

type Handlers struct {
	logger *logger.Logger
	repos  *repository.Repositories
}

func New(logger *logger.Logger, repos *repository.Repositories) *Handlers {
	return &Handlers{
		logger: logger,
		repos:  repos,
	}
}

// internal/api/router.go
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/shinshARK/clickhouse-api/internal/api/handlers"
	customMiddleware "github.com/shinshARK/clickhouse-api/internal/api/middleware"
	"github.com/shinshARK/clickhouse-api/internal/storage/repository"
	pkgLogger "github.com/shinshARK/clickhouse-api/pkg/logger"
)

func SetupRoutes(app *fiber.App, logger *pkgLogger.Logger, repos *repository.Repositories) {
	// Middleware
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New())
	app.Use(customMiddleware.LoggingMiddleware(logger))

	// Initialize handlers
	h := handlers.New(logger, repos)

	// Health route
	app.Get("/health", h.HealthCheck)

	// API routes
	api := app.Group("/api/v1")

	// Users endpoints
	users := api.Group("/users")
	users.Post("/", h.CreateUser)
	users.Get("/", h.GetUsers)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}

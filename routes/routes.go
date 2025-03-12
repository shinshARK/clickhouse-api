package routes

import (
	"crud-fiber-clickhouse/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures API endpoints.
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users", handlers.CreateUserHandler)
	api.Get("/users", handlers.ListUsersHandler)
	api.Get("/users/:id", handlers.GetUserHandler)
	api.Put("/users/:id", handlers.UpdateUserHandler)
	api.Delete("/users/:id", handlers.DeleteUserHandler)
}

package main

import (
	"log"

	"crud-fiber-clickhouse/config"
	"crud-fiber-clickhouse/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize ClickHouse connection
	config.InitClickHouse()

	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// cmd/api/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shinshARK/clickhouse-api/internal/api/router"
	"github.com/shinshARK/clickhouse-api/internal/config"
	"github.com/shinshARK/clickhouse-api/internal/storage/clickhouse"
	"github.com/shinshARK/clickhouse-api/pkg/logger"
)

func main() {
	// Initialize logger
	l := logger.New("info")
	l.Info("Starting API server")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		l.Fatal("Failed to load configuration", "error", err)
	}

	// Connect to ClickHouse
	db, err := clickhouse.Connect(cfg.Database)
	if err != nil {
		l.Fatal("Failed to connect to ClickHouse", "error", err)
	}
	defer db.Close()

	// Initialize repositories
	repos := clickhouse.NewRepositories(db)

	// Setup Fiber app with router
	app := fiber.New(fiber.Config{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		AppName:      "ClickHouse Fiber API",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	// Setup router with handlers
	router.SetupRoutes(app, l, repos)

	// Start server in a goroutine
	go func() {
		l.Info("Server listening", "port", cfg.Server.Port)
		if err := app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
			l.Fatal("Failed to start server", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		l.Fatal("Server forced to shutdown", "error", err)
	}

	l.Info("Server exited properly")
}

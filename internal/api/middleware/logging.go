// internal/api/middleware/logging.go
package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shinshARK/clickhouse-api/pkg/logger"
)

func LoggingMiddleware(logger *logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Log after response
		logger.Info("Request completed",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration", time.Since(start),
			"ip", c.IP(),
			"request_id", c.GetRespHeader("X-Request-ID"),
		)

		return err
	}
}

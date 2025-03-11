// internal/api/handlers/health.go
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "healthy"})
}

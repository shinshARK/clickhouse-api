// internal/api/handlers/user.go
package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shinshARK/clickhouse-api/internal/models"
)

func (h *Handlers) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		h.logger.Error("Failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := h.repos.User.Create(c.Context(), &user); err != nil {
		h.logger.Error("Failed to create user", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *Handlers) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user ID"})
	}

	user, err := h.repos.User.GetByID(c.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get user", "error", err, "id", id)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handlers) GetUsers(c *fiber.Ctx) error {
	limit := 10
	offset := 0

	if limitParam := c.Query("limit"); limitParam != "" {
		if lim, err := strconv.Atoi(limitParam); err == nil && lim > 0 {
			limit = lim
		}
	}

	if offsetParam := c.Query("offset"); offsetParam != "" {
		if off, err := strconv.Atoi(offsetParam); err == nil && off >= 0 {
			offset = off
		}
	}

	users, err := h.repos.User.GetAll(c.Context(), limit, offset)
	if err != nil {
		h.logger.Error("Failed to get users", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get users"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *Handlers) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user ID"})
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		h.logger.Error("Failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Ensure ID in route matches body
	user.ID = id

	if err := h.repos.User.Update(c.Context(), &user); err != nil {
		h.logger.Error("Failed to update user", "error", err, "id", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handlers) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing user ID"})
	}

	if err := h.repos.User.Delete(c.Context(), id); err != nil {
		h.logger.Error("Failed to delete user", "error", err, "id", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": "success"})
}

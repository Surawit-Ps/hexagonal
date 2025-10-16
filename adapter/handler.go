package adapter

import (
	"github.com/gofiber/fiber/v2"
	"hexagonal/core"
)

type Handler struct {
	svc *core.Service
}

func NewHandler(s *core.Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) DeleteEducation(c *fiber.Ctx) error {
	userId := c.Params("id")
	eduId := c.Params("eduId")

	err := h.svc.DeleteEducation(userId, eduId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}	
	return c.JSON(fiber.Map{"message": "education deleted"})
}

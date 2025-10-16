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

func (h *Handler) AddEducation(c *fiber.Ctx) error {
	userId := c.Params("id")
	var edu core.Education
	if err := c.BodyParser(&edu); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.AddEducation(userId, &edu); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "education added"})
}

func (h *Handler) UpdateEducation(c *fiber.Ctx) error {
	userId := c.Params("id")
	eduId := c.Params("eduId")
	var edu core.Education
	if err := c.BodyParser(&edu); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.UpdateEducation(userId, eduId, &edu); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "education updated"})
}

// WorkExp
func (h *Handler) AddWorkExp(c *fiber.Ctx) error {
	userId := c.Params("id")
	var work core.WorkExperience
	if err := c.BodyParser(&work); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.AddWorkExp(userId, &work); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "work experience added"})
}

func (h *Handler) UpdateWorkExp(c *fiber.Ctx) error {
	userId := c.Params("id")
	workId := c.Params("workId")
	var work core.WorkExperience
	if err := c.BodyParser(&work); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.UpdateWorkExp(userId, workId, &work); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "work experience updated"})
}

func (h *Handler) DeleteWorkExp(c *fiber.Ctx) error {
	userId := c.Params("id")
	workId := c.Params("workId")
	if err := h.svc.DeleteWorkExp(userId, workId); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "work experience deleted"})
}

// Project
func (h *Handler) AddProject(c *fiber.Ctx) error {
	userId := c.Params("id")
	workId := c.Params("workId")
	var proj core.Project
	if err := c.BodyParser(&proj); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.AddProject(userId, workId, &proj); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "project added"})
}

func (h *Handler) UpdateProject(c *fiber.Ctx) error {
	userId := c.Params("id")
	workId := c.Params("workId")
	projId := c.Params("projId")
	var proj core.Project
	if err := c.BodyParser(&proj); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.UpdateProject(userId, workId, projId, &proj); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "project updated"})
}

func (h *Handler) DeleteProject(c *fiber.Ctx) error {
	userId := c.Params("id")
	workId := c.Params("workId")
	projId := c.Params("projId")
	if err := h.svc.DeleteProject(userId, workId, projId); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "project deleted"})
}

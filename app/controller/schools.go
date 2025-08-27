package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/app/repository"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

func CreateSchool(c *fiber.Ctx) error {
	school := &model.School{}
	if err := c.BodyParser(school); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"msg":    "Invalid request body",
			"status": fiber.StatusBadRequest,
		})
	}

	schoolRepo := repository.NewSchoolRepo(database.GetDB())
	err := schoolRepo.Create(school)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"msg": "Can't create school",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"msg":    "Create school successfully",
		"status": fiber.StatusCreated,
		"data":   school,
	})
}

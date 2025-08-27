package controller

import (
	"github.com/gofiber/fiber/v2"
)

func CreateVoter(c *fiber.Ctx) error {
	// Logic to create an Voter
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Voter created successfully",
	})
}

func GetVoters(c *fiber.Ctx) error {
	// Logic to get all Voters
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "List of Voters",
	})
}

func DeleteVoter(c *fiber.Ctx) error {
	// Logic to delete an Voter
	id := c.Params("id")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Voter deleted successfully",
		"id":  id,
	})
}

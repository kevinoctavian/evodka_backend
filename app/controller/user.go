package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	// Logic to get all users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "List of users",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// Logic to update a user
	id := c.Params("id")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "User updated successfully",
		"id":  id,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	// Logic to delete a user
	id := c.Params("id")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "User deleted successfully",
		"id":  id,
	})
}

package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/gofiber/fiber/v2"
)

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	err := data_access.TodoDelete(id)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}

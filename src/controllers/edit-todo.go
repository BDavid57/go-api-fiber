package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

func EditTodo(c *fiber.Ctx) error {
	var todo dto.Todo
	err := c.BodyParser(&todo)
	id := c.Params("id")

	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data_access.TodoEdit(id, todo)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}

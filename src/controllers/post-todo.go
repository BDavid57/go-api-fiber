package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

func PostTodo(c *fiber.Ctx) error {
	var todo dto.Todo
	err := c.BodyParser(&todo)

	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	todo, _ = data_access.TodoCreate(todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

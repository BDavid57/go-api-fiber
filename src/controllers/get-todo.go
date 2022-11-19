package controllers

import (
	"log"

	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/gofiber/fiber/v2"
)

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := data_access.TodoGet(id)

	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

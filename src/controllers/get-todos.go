package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	todos, err := data_access.TodosGet()

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	response := dto.GetAllTodosResponse {
		Data: todos,
		Total: len(todos),
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Get all from db
func GetTodos(c *fiber.Ctx) error {
	todos, err := data_access.TodosGet()

	response := dto.GetAllTodosResponse {
		Data: todos,
		Total: len(todos),
	}

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(response)
}

// Get all from hardcoded data
func GetTweets(c *fiber.Ctx) error {
	response := dto.GetAllTweetsResponse {
		Data: data.Tweets,
		Total: len(data.Tweets),
	}

	return c.JSON(response)
}

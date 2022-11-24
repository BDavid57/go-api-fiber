package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Get all from db
func GetTweets(c *fiber.Ctx) error {
	tweets, err := data_access.TweetsGet()

	response := dto.GetAllTweetsResponse {
		Data: tweets,
		Total: len(tweets),
	}

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(response)
}

// Get all from hardcoded data
func GetTodos(c *fiber.Ctx) error {
	response := dto.GetAllTodosResponse {
		Data: data.Todos,
		Total: len(data.Todos),
	}

	return c.JSON(response)
}

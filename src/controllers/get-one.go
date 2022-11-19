package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/gofiber/fiber/v2"
)

// Get one from db
func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := data_access.TodoGet(id)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

// Get one from hardcoded data
func GetTweetById(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, item := range data.Tweets {
		if item.ID == id {
			return c.JSON(item)
			
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Tweet not found",
	})
}

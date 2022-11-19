package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Add one to db
func PostTodo(c *fiber.Ctx) error {
	var todo dto.Todo
	err := c.BodyParser(&todo)

	data_access.TodoCreate(todo)
	
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

// Add one to hardcoded data
func PostTweet(c *fiber.Ctx) error {
	var newTweet dto.Tweet

	if err := c.BodyParser(&newTweet); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Couldn't create tweet",
		})
	}

	data.Tweets = append(data.Tweets, newTweet)
	return c.Status(fiber.StatusCreated).JSON(newTweet)
}

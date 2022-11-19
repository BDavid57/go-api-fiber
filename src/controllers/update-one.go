package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Edit one from db
func EditTodo(c *fiber.Ctx) error {
	var todo dto.Todo
	err := c.BodyParser(&todo)
	id := c.Params("id")

	data_access.TodoEdit(id, todo)
	
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}

// Edit one from hardcoded data
func EditTweet(c *fiber.Ctx) error {
	id := c.Params("id")
	newSlice := []dto.Tweet{}

	var newTweet dto.Tweet

	if err := c.BodyParser(&newTweet); err != nil {
		return err
	}

	for _, tweet := range data.Tweets {
		if tweet.ID != id {
			newSlice = append(newSlice, tweet)
			continue
		}
		if tweet.ID == id {
			newSlice = append(newSlice, newTweet)
		}
	}

	data.Tweets = newSlice
	return c.JSON(newTweet)
}

package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/data_access"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Delete from db
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

// Delete from hardcoded data
func DeleteTweet(c *fiber.Ctx) error {
	id := c.Params("id")
	newSlice := []dto.Tweet{}

	for _, tweet := range data.Tweets {
		if tweet.ID != id {
			newSlice = append(newSlice, tweet)
		}
	}

	if len(newSlice) == len(data.Tweets) {
		return fiber.NewError(fiber.StatusNotFound, "Tweet not found!")
	}

	data.Tweets = newSlice
	return c.JSON(fiber.Map{
		"message": "Tweet deleted successfully",
	})
}

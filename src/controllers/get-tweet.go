package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/gofiber/fiber/v2"
)

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

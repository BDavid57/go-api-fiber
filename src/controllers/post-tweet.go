package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

func PostTweet(c *fiber.Ctx) error {
	var newTweet dto.Tweet

	if err := c.BodyParser(&newTweet); err != nil {
		return err
	}

	data.Tweets = append(data.Tweets, newTweet)
	return c.JSON(newTweet)
}

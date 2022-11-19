package controllers

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

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
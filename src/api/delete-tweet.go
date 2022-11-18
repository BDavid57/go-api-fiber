package api

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)


func DeleteTweet(c *fiber.Ctx) error {
	id := c.Params("id")
	newSlice := []dto.Tweet{}

	message := dto.NewMessage("Tweet deleted succesfully")

	for _, tweet := range data.Tweets {
		if tweet.ID != id {
			newSlice = append(newSlice, tweet)
		}
	}

	if len(newSlice) == len(data.Tweets) {
		return fiber.NewError(fiber.StatusNotFound, "Tweet not found!")
	}

	data.Tweets = newSlice
	return c.JSON(message)
}
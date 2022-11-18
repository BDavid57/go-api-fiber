package api

import (
	"github.com/BDavid57/go-api-fiber/src/data"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

func GetTweets(c *fiber.Ctx) error {
	response := dto.GetAllTweetsResponse {
		Data: data.Tweets,
		Total: len(data.Tweets),
	}

	return c.JSON(response)
}

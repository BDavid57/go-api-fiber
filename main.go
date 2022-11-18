package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/BDavid57/go-api-fiber/src/api"
)

func main() {
	app := fiber.New()

	tweetsApi := app.Group("/api")
		
	tweetsApi.Get("/tweets", api.GetTweets)
	tweetsApi.Get("/tweets/:id", api.GetTweetById)
	tweetsApi.Post("/tweets", api.PostTweet)
	tweetsApi.Delete("/tweets/:id", api.DeleteTweet)
	tweetsApi.Put("/tweets/:id", api.EditTweet)

	app.Listen("localhost:5000")
}

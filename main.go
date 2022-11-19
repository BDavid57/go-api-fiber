package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/BDavid57/go-api-fiber/src/controllers"
)

func main() {
	app := fiber.New()

	appApi := app.Group("/api")

	// comes from db
	appApi.Post("/todo", controllers.PostTodo)
	appApi.Get("/todo/:id", controllers.GetTodo)
	appApi.Get("/todo", controllers.GetTodos)
	appApi.Delete("/todo/:id", controllers.DeleteTodo)
		
	// comes from hardcoded data
	appApi.Get("/tweets", controllers.GetTweets)
	appApi.Get("/tweets/:id", controllers.GetTweetById)
	appApi.Post("/tweets", controllers.PostTweet)
	appApi.Delete("/tweets/:id", controllers.DeleteTweet)
	appApi.Put("/tweets/:id", controllers.EditTweet)

	app.Listen("localhost:5000")
}

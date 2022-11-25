package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/BDavid57/go-api-fiber/src/controllers"
	"github.com/BDavid57/go-api-fiber/src/db"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	appApi := app.Group("/api")

	// comes from db
	appApi.Get("/tweets", controllers.GetTweets)
	appApi.Get("/tweets/:id", controllers.GetTweetById)
	appApi.Post("/tweets", controllers.PostTweet)
	appApi.Delete("/tweets/:id", controllers.DeleteTweet)
	appApi.Put("/tweets/:id", controllers.EditTweet)
	
	// comes from hardcoded data
	appApi.Post("/todo", controllers.PostTodo)
	appApi.Get("/todo/:id", controllers.GetTodo)
	appApi.Get("/todo", controllers.GetTodos)
	appApi.Delete("/todo/:id", controllers.DeleteTodo)
	appApi.Put("/todo/:id", controllers.EditTodo)
		
	db.ConnectToDb()
	app.Listen("localhost:5000")
}

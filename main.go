package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/BDavid57/go-api-fiber/src/controllers"
	"github.com/BDavid57/go-api-fiber/src/db"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	
	appApi := app.Group("/api")

	appApi.Get("/tweets", controllers.GetTweets)
	appApi.Get("/tweets/:id", controllers.GetTweetById)
	appApi.Post("/tweets", controllers.PostTweet)
	appApi.Delete("/tweets/:id", controllers.DeleteTweet)
	appApi.Put("/tweets/:id", controllers.EditTweet)
		
	db.ConnectToDb()
	app.Listen("localhost:5000")
}

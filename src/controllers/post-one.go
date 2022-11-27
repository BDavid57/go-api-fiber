package controllers

import (
	"context"
	"log"
	"time"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Add one to db
func PostTweet(c *fiber.Ctx) error {
	twitterCloneCollection := db.DB.Collection("twitter_clone")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	tweet := new(dto.Tweet)

	
	if err := c.BodyParser(tweet); err != nil {
        log.Println(err)
        return c.Status(400).JSON(fiber.Map{
            "message": "Failed to parse body",
            "error":   err,
        })
    }

    result, err := twitterCloneCollection.InsertOne(ctx, tweet)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Tweet failed to insert",
            "error":   err,
        })
    }

    var foundTweet dto.Tweet
    findResult := twitterCloneCollection.FindOne(ctx, bson.M{"_id": result.InsertedID})
    if err := findResult.Err(); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Tweet Not found",
            "error":   err,
        })
    }

    err = findResult.Decode(&foundTweet)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Tweet Not found",
            "error":   err,
        })
    }

    return c.Status(fiber.StatusCreated).JSON(foundTweet)
}

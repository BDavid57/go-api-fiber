package controllers

import (
	"context"
	"time"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get all from db
func GetTweets(c *fiber.Ctx) error {
	twitterCloneCollection := db.DB.Collection("twitter_clone")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    var tweets []dto.Tweet

    filter := bson.M{}
    findOptions := options.Find()

    total, _ := twitterCloneCollection.CountDocuments(ctx, filter)

    cursor, err := twitterCloneCollection.Find(ctx, filter, findOptions)

    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Tweets Not found",
            "error":   err,
        })
    }

    for cursor.Next(ctx) {
        var tweet dto.Tweet
        cursor.Decode(&tweet)
        tweets = append(tweets, tweet)
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "data":      tweets,
        "total":     total,
    })
}

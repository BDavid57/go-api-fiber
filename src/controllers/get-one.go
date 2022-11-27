package controllers

import (
	"context"
	"time"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get one from db
func GetTweetById(c *fiber.Ctx) error {
	twitterCloneCollection := db.DB.Collection("twitter_clone")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    var tweet dto.Tweet
    objId, err := primitive.ObjectIDFromHex(c.Params("id"))
    findResult := twitterCloneCollection.FindOne(ctx, bson.M{"_id": objId})
    if err := findResult.Err(); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Tweet Not found",
            "error":   err,
        })
    }

    err = findResult.Decode(&tweet)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Tweet Not found",
            "error":   err,
        })
    }

    return c.Status(fiber.StatusOK).JSON(tweet)
}

package controllers

import (
	"context"
	"time"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete from db
func DeleteTweet(c *fiber.Ctx) error {
	twitterCloneCollection := db.DB.Collection("twitter_clone")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    objId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Tweet not found",
            "error":   err,
        })
    }
    _, err = twitterCloneCollection.DeleteOne(ctx, bson.M{"_id": objId})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Tweet failed to delete",
            "error":   err,
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Tweet deleted successfully",
    })
}

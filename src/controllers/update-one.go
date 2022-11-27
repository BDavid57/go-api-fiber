package controllers

import (
	"context"
	"log"
	"time"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Edit one from db
func EditTweet(c *fiber.Ctx) error {
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

    objId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Tweet not found",
            "error":   err,
        })
    }

    update := bson.M{
        "$set": tweet,
    }
    _, err = twitterCloneCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Tweet failed to update",
            "error":   err.Error(),
        })
    }
	
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Tweet updated successfully",
    })
}

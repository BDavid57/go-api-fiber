package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TweetDelete(tweetId string) error {
	twitterCloneCollection := db.DB.Collection("twitter_clone")

	objectId, err := primitive.ObjectIDFromHex(tweetId)
	if err != nil {
		return err
	}

	_, err = twitterCloneCollection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}

package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TweetGet(tweetId string) (dto.Tweet, error){
	twitterCloneCollection := db.DB.Collection("twitter_clone")
	var tweet dto.Tweet

	objectId, err := primitive.ObjectIDFromHex(tweetId)
	if err != nil {
		return tweet, err
	}

	err = twitterCloneCollection.FindOne(context.Background(), bson.M{"_id":objectId}).Decode(&tweet)
	if err != nil {
		return tweet, err
	}

	return tweet, nil
}

package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func TweetsGet() ([]dto.Tweet, error){
	twitterCloneCollection := db.DB.Collection("twitter_clone")
	var tweets []dto.Tweet

	res, err := twitterCloneCollection.Find(context.Background(), bson.D{})

	if res.All(context.Background(), &tweets); err != nil {
		return tweets, err
	}

	return tweets, nil
}

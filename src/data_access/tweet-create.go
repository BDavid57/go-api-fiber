package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
)

func TweetCreate(tweet dto.Tweet) (dto.Tweet, error) {
	twitterCloneCollection := db.DB.Collection("twitter_clone")
	
	_, err := twitterCloneCollection.InsertOne(context.Background(), tweet)

	if err != nil {
		return dto.Tweet{}, err
	}

	return tweet, nil
}

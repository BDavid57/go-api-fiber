package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TodoDelete(todoId string) error {
	todoCollection := db.DB.Collection("todo")

	objectId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return err
	}

	_, err = todoCollection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		return err
	}


	return nil
}

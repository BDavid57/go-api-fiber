package data_access

import (
	"context"
	"log"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TodoEdit(todoId string, todo dto.Todo) error {
	todoCollection := db.DB.Collection("todo")

	objectId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		log.Fatal(err)
		return err
	}

	update := bson.M{
        "$set": todo,
    }
	_, err = todoCollection.UpdateOne(context.Background(), bson.M{"_id": objectId}, update)
	if err != nil {
		log.Fatal(err)
		return err
	}


	return nil
}

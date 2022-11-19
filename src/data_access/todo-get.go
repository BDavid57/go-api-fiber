package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TodoGet(todoId string) (dto.Todo, error){
	todoCollection := db.DB.Collection("todo")
	var todo dto.Todo

	objectId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return todo, err
	}

	err = todoCollection.FindOne(context.Background(), bson.M{"_id":objectId}).Decode(&todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

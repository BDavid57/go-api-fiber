package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func TodosGet() ([]dto.Todo, error){
	todoCollection := db.DB.Collection("todo")
	var todos []dto.Todo

	res, err := todoCollection.Find(context.Background(), bson.D{})

	if res.All(context.Background(), &todos); err != nil {
		return todos, err
	}

	return todos, nil
}

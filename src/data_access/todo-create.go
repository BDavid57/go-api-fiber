package data_access

import (
	"context"

	"github.com/BDavid57/go-api-fiber/src/db"
	"github.com/BDavid57/go-api-fiber/src/dto"
)

func TodoCreate(todo dto.Todo) (dto.Todo, error) {
	todoCollection := db.DB.Collection("todo")
	
	_, err := todoCollection.InsertOne(context.Background(), todo)

	if err != nil {
		return dto.Todo{}, err
	}

	return todo, nil
}

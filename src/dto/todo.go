package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

type GetAllTodosResponse struct {
	Data  []Todo `json:"data"`
	Total int    `json:"total"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo struct to describe todo object.
type Todo struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omitempty"`
}

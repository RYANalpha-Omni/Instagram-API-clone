package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Email  string             `json:"title" bson:"title,omitempty"`
	Password *User            `json:"userr" bson:"user,omitempty"`
}

type Post struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption string            `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Image   string            `json:"title" bson:"title,omitempty"`
	Timestamp *Time           `json:"$now()" bson:"author,omitempty"`
}
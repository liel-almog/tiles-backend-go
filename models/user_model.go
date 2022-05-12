package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name     string             `json:"name" bson:"name" validate:"required"`
	Email    string             `json:"email" bson:"email" validate:"required"`
	Password string             `json:"password" bson:"password" validate:"required"`
	Role     string             `json:"role" bson:"role"`
}

type UserDTO struct {
	Email string             `json:"email" bson:"email"`
	Name  string             `json:"name" bson:"name"`
	Role  string             `json:"role" bson:"role"`
	Id    primitive.ObjectID `json:"id" bson:"_id"`
}

type UpdateRole struct {
	Role string             `json:"role" bson:"role" validate:"required"`
	Id   primitive.ObjectID `json:"id" bson:"id" validate:"required"`
}

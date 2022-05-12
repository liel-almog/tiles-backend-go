package res

import (
	"tiles/tiles-backend-go/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRes struct {
	User    models.UserDTO `json:"user"`
	Token   string         `json:"token"`
	Message string         `json:"message"`
	Status  int            `json:"status"`
}

type RegisterRes struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    mongo.InsertOneResult `json:"data"`
}

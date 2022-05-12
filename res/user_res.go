package res

import "tiles/tiles-backend-go/models"

type UserRes struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    models.UserDTO `json:"data"`
}

type UsersRes struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []models.UserDTO `json:"data"`
}

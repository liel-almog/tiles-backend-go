package res

import "tiles/tiles-backend-go/models"

type TilesRes struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Tile `json:"data"`
}

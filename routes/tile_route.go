package routes

import (
	"tiles/tiles-backend-go/controllers"

	"github.com/gin-gonic/gin"
)

func TileRoute(router *gin.Engine) {
	router.GET("/tile", controllers.GetAllTiles())

	router.PATCH("/tile", controllers.UpdateTiles())
}

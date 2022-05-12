package main

import (
	"tiles/tiles-backend-go/configs"
	"tiles/tiles-backend-go/middleware"
	"tiles/tiles-backend-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"http://localhost:3000/"})

	router.Use(middleware.Cors())

	configs.ConnectDB()

	routes.UserRoute(router)
	routes.AuthRoute(router)
	routes.TileRoute(router)

	router.Run("localhost:8080")
}

package routes

import (
	"tiles/tiles-backend-go/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/auth/login", controllers.GetAUser())

	router.POST("/auth/signup", controllers.CreateUser())
}

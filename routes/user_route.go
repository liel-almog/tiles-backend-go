package routes

import (
	"tiles/tiles-backend-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user", controllers.GetAllUsers())

	router.GET("/user/role/:role", controllers.GetUsersByRole())

	router.PATCH("/user/role", controllers.UpdateRoles())
}

package routes

import "github.com/gin-gonic/gin"

func UserRoute(router *gin.Engine) {
	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/user/role/:role", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.PATCH("/user/role", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}

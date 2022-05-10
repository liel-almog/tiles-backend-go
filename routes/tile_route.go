package routes

import "github.com/gin-gonic/gin"

func TileRoute(router *gin.Engine) {
	router.GET("/tile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.PATCH("/tile", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}

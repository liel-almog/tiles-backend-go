package routes

import "github.com/gin-gonic/gin"

func AuthRoute(router *gin.Engine) {
	router.POST("/auth/signup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.POST("/auth/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}

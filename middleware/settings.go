package middleware

import (
	"tiles/tiles-backend-go/configs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	clientURL := configs.GetEnv("CLIENT_URL")

	return cors.New(cors.Config{
		AllowOrigins:     []string{clientURL},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	})
}

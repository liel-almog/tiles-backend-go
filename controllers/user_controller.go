package controllers

import (
	"context"
	"net/http"
	"tiles/tiles-backend-go/models"
	"tiles/tiles-backend-go/res"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout((context.Background()), 10*time.Second)

		var users []models.UserDTO
		defer cancel()

		results, err := userCollection.Find(ctx, bson.D{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, res.ErrorRes{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"error": err.Error()},
			})
			return
		}

		for results.Next(ctx) {
			var user models.UserDTO
			if err := results.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, res.ErrorRes{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"error": err.Error()}})
				return
			}

			users = append(users, user)
		}

		c.JSON(http.StatusOK, res.UsersRes{
			Status:  http.StatusOK,
			Message: "success",
			Data:    users,
		})
	}
}

func GetUsersByRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout((context.Background()), 10*time.Second)

		var users []models.UserDTO
		defer cancel()

		role := c.Param("role")
		results, err := userCollection.Find(ctx, bson.M{"role": role})

		if err != nil {
			c.JSON(http.StatusInternalServerError, res.ErrorRes{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		for results.Next(ctx) {
			var user models.UserDTO
			if err := results.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, res.ErrorRes{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"error": err.Error()},
				})

				return
			}

			users = append(users, user)
		}

		c.JSON(http.StatusOK, res.UsersRes{
			Status:  http.StatusOK,
			Message: "success",
			Data:    users,
		})
	}
}

func UpdateRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout((context.Background()), 10*time.Second)

		var updateRole []models.UpdateRole
		defer cancel()

		if err := c.BindJSON(&updateRole); err != nil {
			c.JSON(http.StatusBadRequest, res.ErrorRes{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		// if validationErr := validate.Struct(updateRole); validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, res.ErrorRes{
		// 		Status:  http.StatusBadRequest,
		// 		Message: "Validation Error",
		// 		Data:    map[string]interface{}{"error": validationErr.Error()},
		// 	})

		// 	return
		// }

		var models []mongo.WriteModel
		for _, user := range updateRole {
			models = append(models, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": user.Id}).
				SetUpdate(bson.M{"$set": bson.M{"role": user.Role}}))
		}
		opts := options.BulkWrite().SetOrdered(false)

		results, err := userCollection.BulkWrite(ctx, models, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, res.ErrorRes{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    results,
		})
	}
}

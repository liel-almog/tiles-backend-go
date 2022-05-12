package controllers

import (
	"context"
	"net/http"
	"tiles/tiles-backend-go/configs"
	"tiles/tiles-backend-go/models"
	"tiles/tiles-backend-go/res"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tileCollection *mongo.Collection = configs.GetCollection(configs.DB, "tiles")

func GetAllTiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		tiles := make([]models.Tile, 0)
		defer cancel()

		results, err := tileCollection.Find(ctx, bson.D{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, res.ErrorRes{
				Message: "Internal server error",
				Status:  http.StatusInternalServerError,
				Data:    map[string]interface{}{"data": err.Error()},
			})

			return
		}

		for results.Next(ctx) {
			var tile models.Tile

			if err := results.Decode(&tile); err != nil {
				c.JSON(http.StatusInternalServerError, res.ErrorRes{
					Message: "Internal server error",
					Status:  http.StatusInternalServerError,
					Data:    map[string]interface{}{"data": err.Error()},
				})

				return
			}

			tiles = append(tiles, tile)
		}

		c.JSON(http.StatusOK, res.TilesRes{
			Message: "success",
			Status:  http.StatusOK,
			Data:    tiles,
		})
	}
}

func UpdateTiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var updateTiles models.UpdateTiles
		defer cancel()

		if err := c.BindJSON(&updateTiles); err != nil {
			c.JSON(http.StatusBadRequest, res.ErrorRes{
				Message: "Bad request",
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"data": err.Error()},
			})

			return
		}

		var models []mongo.WriteModel

		for _, tile := range updateTiles.Added {

			models = append(models, mongo.NewInsertOneModel().SetDocument(tile))
		}

		for _, objId := range updateTiles.Deleted {
			models = append(models, mongo.NewDeleteOneModel().SetFilter(bson.M{"_id": objId}))
		}

		for _, tile := range updateTiles.Changed {
			models = append(models, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": tile.Id}).
				SetUpdate(bson.M{"$set": bson.M{"color": tile.Color}}))
		}

		opts := options.BulkWrite().SetOrdered(false)

		results, err := tileCollection.BulkWrite(ctx, models, opts)
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

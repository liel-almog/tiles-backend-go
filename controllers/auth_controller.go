package controllers

import (
	"context"
	"net/http"
	"tiles/tiles-backend-go/configs"
	"tiles/tiles-backend-go/models"
	"tiles/tiles-backend-go/res"
	"tiles/tiles-backend-go/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout((context.Background()), 10*time.Second)

		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, res.ErrorRes{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, res.ErrorRes{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"error": validationErr.Error()},
			})

			return
		}

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		newUser := models.Signup{
			Role:     "Viewer",
			Name:     user.Name,
			Email:    user.Email,
			Password: string(encryptedPassword),
		}

		result, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, res.ErrorRes{
				Message: "Internal server error",
				Status:  http.StatusInternalServerError,
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		c.JSON(http.StatusOK, res.RegisterRes{
			Message: "User created",
			Status:  http.StatusCreated,
			Data:    *result,
		})
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout((context.Background()), 10*time.Second)

		var login models.Login
		defer cancel()

		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, res.ErrorRes{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"error": err.Error()},
			})
			return
		}

		if validationErr := validate.Struct(&login); validationErr != nil {
			c.JSON(http.StatusBadRequest, res.ErrorRes{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"error": validationErr.Error()},
			})

			return
		}

		var userFound models.User
		err := userCollection.FindOne(ctx, bson.M{"email": login.Email}).Decode(&userFound)

		if err != nil {
			c.JSON(http.StatusNotFound, res.ErrorRes{
				Message: "Email or Password is incorrect",
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(login.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, res.ErrorRes{
				Message: "Email or Password is incorrect",
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		loginUser := models.UserDTO{
			Email: userFound.Email,
			Name:  userFound.Name,
			Role:  userFound.Role,
			Id:    userFound.Id,
		}

		token, err := util.GenerateJWT(loginUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, res.ErrorRes{
				Message: "Internal server error",
				Status:  http.StatusInternalServerError,
				Data:    map[string]interface{}{"error": err.Error()},
			})

			return
		}

		c.JSON(http.StatusOK, res.LoginRes{
			Message: "User found",
			Token:   token,
			Status:  http.StatusOK,
			User:    loginUser,
		})

	}
}

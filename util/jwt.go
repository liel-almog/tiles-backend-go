package util

import (
	"fmt"
	"tiles/tiles-backend-go/configs"
	"tiles/tiles-backend-go/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.UserDTO) (string, error) {
	mySigningKey := []byte(configs.GetEnv("JWT_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["role"] = user.Role
	claims["id"] = user.Id

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

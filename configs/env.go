package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoUri() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

func GetEnv(env string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envVal, isEnv := os.LookupEnv(env)
	if !isEnv {
		log.Fatal("Environment variable not found")
	}

	return envVal
}

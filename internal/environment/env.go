package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(envFiles ...string) {
	err := godotenv.Load(envFiles...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}

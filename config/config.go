package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv(value string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(value)
}

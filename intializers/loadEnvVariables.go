package intializers

import (
	"log"

	"github.com/joho/godotenv"
)

// function LoadEnvironmentVariables() that loads environment variables from a .env file using the godotenv library
func LoadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

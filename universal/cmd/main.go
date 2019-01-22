package main

import (
	"log"
	"os"

	"github.com/calazans10/go-structure-examples/universal/pkg/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	server := api.Server{}
	server.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	server.Run(
		os.Getenv("APP_ENV"),
		os.Getenv("APP_PORT"),
	)
}

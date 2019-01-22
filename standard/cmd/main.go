package main

import (
	"log"
	"os"

	"github.com/calazans10/go-structure-examples/standard/pkg/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("APP_DB_USERNAME")
	dbPass := os.Getenv("APP_DB_PASSWORD")
	dbName := os.Getenv("APP_DB_NAME")
	appEnv := os.Getenv("APP_ENV")
	appPort := os.Getenv("APP_PORT")

	server := api.Server{}
	server.Initialize(dbUser, dbPass, dbName)
	server.Run(appEnv, appPort)
}

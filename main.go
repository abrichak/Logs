package main

import (
	"log"
	"metrics-service/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := server.NewServer()

	err = app.Start(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Port already used")
	}
}

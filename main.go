package main

import (
	"log"
	"metrics-service/server"
	"metrics-service/server/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := server.NewServer()

	routes.ConfigureRoutes(app)
	err = app.Start(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Port already used")
	}
}

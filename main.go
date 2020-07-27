package main

import (
	"fmt"
	"log"
	"metrics-service/docs"
	"metrics-service/server"
	"metrics-service/server/routes"
	"os"

	"github.com/joho/godotenv"
)

// @title Metrics Service API
// @version 1.0
// @description This is an API for metrics service

// @contact.name NIX Solutions
// @contact.url https://www.nixsolutions.com/
// @contact.email ask@nixsolutions.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("EXPOSE_PORT"))

	app := server.NewServer()

	routes.ConfigureRoutes(app)
	err = app.Start(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Port already used")
	}
}

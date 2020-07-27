package helpers

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"metrics-service/server"
	"metrics-service/server/routes"
)

func NewServer() *server.Server {
	loadEnv()

	s := &server.Server{
		Echo:  echo.New(),
	}

	routes.ConfigureRoutes(s)

	return s
}

func loadEnv() {
	if err := godotenv.Load("../.env.testing"); err != nil {
		log.Fatal("Error loading .env file for testing")
	}
}
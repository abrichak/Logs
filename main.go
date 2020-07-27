package main

import (
	"fmt"
	"log"
	"metrics-service/docs"
	"metrics-service/server"
	"metrics-service/server/routes"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
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

	var g errgroup.Group

	appLogs := server.NewServer()
	routes.ConfigureLogsRoutes(appLogs)

	appMetrics := server.NewServer()
	routes.ConfigureMetricsRoutes(appMetrics)

	g.Go(func() error {
		return appLogs.Start(os.Getenv("PORT"))
	})

	g.Go(func() error {
		return appMetrics.Start(os.Getenv("PORT_METRICS"))
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

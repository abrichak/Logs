package main

import (
	"fmt"
	"log"
	"metrics-service/docs"
	"metrics-service/server"
	"metrics-service/server/routes"
	"metrics-service/server/services"
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

	// We need 2 app instances for different ports, because the framework allows the routing
	// only using 1 port per application instance
	appMetrics := server.NewServer()
	uniqueIPs := routes.ConfigureMetricsRoutes(appMetrics)

	appLogs := server.NewServer()
	routes.ConfigureLogsRoutes(appLogs, uniqueIPs)

	// Delete the IP addresses list on service start according to the requirements
	appLogs.Redis.Del(services.CacheKeyForIps)

	var g errgroup.Group
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

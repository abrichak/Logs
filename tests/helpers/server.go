package helpers

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"metrics-service/server"
	"metrics-service/server/routes"
)

func NewServer() *server.Server {
	loadEnv()

	_, redisClient := InitRedis()
	s := &server.Server{
		Echo:  echo.New(),
		Redis: redisClient,
	}

	uniqueIPsMetric := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "unique_ip_addresses",
	})
	routes.ConfigureLogsRoutes(s, &uniqueIPsMetric)

	return s
}

func loadEnv() {
	if err := godotenv.Load("../.env.testing"); err != nil {
		log.Fatal("Error loading .env file for testing")
	}
}
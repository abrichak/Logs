package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	s "metrics-service/server"
	"metrics-service/server/handlers"
)

func ConfigureLogsRoutes(server *s.Server, uniqueIPsMetric *prometheus.Gauge) {
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// We need to pass the custom Prometheus metric to the handler and further in order to call it's Set() method
	// to set the correct value of unique IP addresses after each "/logs" request
	server.Echo.POST("/logs", handlers.NewLogsHandler(server, uniqueIPsMetric).SaveLogMessage)
}

func ConfigureMetricsRoutes(server *s.Server) *prometheus.Gauge {
	server.Echo.Use(middleware.Logger())

	uniqueIPsMetric := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "unique_ip_addresses",
	})

	server.Echo.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return &uniqueIPsMetric
}
package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	s "metrics-service/server"
	"metrics-service/server/handlers"
)

func ConfigureLogsRoutes(server *s.Server) {
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/logs", handlers.NewLogsHandler(server).SaveLogMessage)
}

func ConfigureMetricsRoutes(server *s.Server) {
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
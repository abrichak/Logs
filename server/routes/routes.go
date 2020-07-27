package routes

import (
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	s "metrics-service/server"
	"metrics-service/server/handlers"
)

func ConfigureRoutes(server *s.Server) {
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/logs", handlers.NewLogsHandler(server).SaveLogMessage)
}
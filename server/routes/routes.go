package routes

import (
	"github.com/labstack/echo/v4/middleware"
	s "metrics-service/server"
)

func ConfigureRoutes(server *s.Server) {
	server.Echo.Use(middleware.Logger())
}
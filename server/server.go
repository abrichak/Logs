package server

import (
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"metrics-service/server/db"
)

type Server struct {
	Echo  *echo.Echo
	Redis *redis.Client
}

func NewServer() *Server {
	return &Server{
		Echo:  echo.New(),
		Redis: db.InitRedis(),
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}

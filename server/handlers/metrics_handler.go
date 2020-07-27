package handlers

import (
	"github.com/labstack/echo/v4"
	s "metrics-service/server"
	"metrics-service/server/responses"
	"metrics-service/server/services"
	"strconv"
)

type MetricsHandler struct {
	server *s.Server
}

func NewMetricsHandler(server *s.Server) *MetricsHandler {
	return &MetricsHandler{server: server}
}

func (handler *MetricsHandler) GetIPsCount(c echo.Context) error {
	count, err := handler.server.Redis.HLen(services.CacheKeyForIps).Result()
	if err != nil {
		panic(err)
	}

	return responses.SuccessResponse(c, strconv.FormatInt(count, 10))
}

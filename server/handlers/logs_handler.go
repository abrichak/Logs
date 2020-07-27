package handlers

import (
	s "metrics-service/server"
	"metrics-service/server/requests"
	"metrics-service/server/responses"
	"metrics-service/server/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

type LogsHandler struct {
	server          *s.Server
	uniqueIPsMetric *prometheus.Gauge
}

func NewLogsHandler(server *s.Server, uniqueIPsMetric *prometheus.Gauge) *LogsHandler {
	return &LogsHandler{
		server: server,
		uniqueIPsMetric: uniqueIPsMetric,
	}
}

// SaveLog godoc
// @Summary Parse the log message and save results
// @Description Parse the log message and save results
// @ID save-log
// @Tags Log Actions
// @Accept json
// @Produce json
// @Param params body requests.SaveLogRequest true "Log message, including timestamp, user's IP, URL address"
// @Success 200 {object} responses.Success
// @Router /logs [post]
func (handler *LogsHandler) SaveLogMessage(c echo.Context) error {
	logRequest := new(requests.SaveLogRequest)
	if err := c.Bind(logRequest); err != nil {
		return err
	}

	logService := services.NewLogService(handler.server, handler.uniqueIPsMetric)
	if err := logService.SaveLogMessage(logRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Server error")
	}

	return responses.SuccessResponse(c, "Log message has been saved")
}

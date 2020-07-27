package services

import (
	"metrics-service/server/requests"
)

type LogService struct {
}

func NewLogService() *LogService {
	return &LogService{}
}

func (service LogService) SaveLogMessage(request *requests.SaveLogRequest) error {
	return nil
}

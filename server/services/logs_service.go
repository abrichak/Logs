package services

import (
	"metrics-service/server"
	"metrics-service/server/requests"
)

const CacheKeyForIps = "user_ip"

type LogService struct {
	Server *server.Server
}

func NewLogService(server *server.Server) *LogService {
	return &LogService{server}
}

func (service LogService) SaveLogMessage(request *requests.SaveLogRequest) error {

	// I have chosen the caching with Redis here because:
	// 1. The issues with memory usage and speed have been already resolved by Redis creators.
	// 2. Redis DB allows horizontal scalability
	// 3. In case we need to add logic to our logs processing (e.g., how many times user requested our site), we can use
	// the "hash" Redis type to save the necessary info
	service.Server.Redis.HSet(CacheKeyForIps, request.IP, "1")

	return nil
}

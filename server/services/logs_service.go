package services

import (
	"github.com/prometheus/client_golang/prometheus"
	"metrics-service/server"
	"metrics-service/server/requests"
)

const CacheKeyForIps = "user_ip"

type LogService struct {
	Server 		    *server.Server
	UniqueIPsMetric *prometheus.Gauge
}

func NewLogService(server *server.Server, uniqueIPsMetric *prometheus.Gauge) *LogService {
	return &LogService{
		server,
		uniqueIPsMetric,
	}
}

func (service LogService) SaveLogMessage(request *requests.SaveLogRequest) error {

	// I have chosen the caching with Redis here because:
	// 1. The issues with memory usage and speed have been already resolved by Redis creators.
	// 2. Redis DB allows horizontal scalability
	// 3. In case we need to add logic to our logs processing (e.g., how many times user requested our site), we can use
	// the "hash" Redis type to save the necessary info
	if err := service.Server.Redis.HSet(CacheKeyForIps, request.IP, "1").Err(); err != nil {
		return err
	}

	return service.renewPromMetric()
}

func (service LogService) renewPromMetric() error {
	count, err := service.Server.Redis.HLen(CacheKeyForIps).Result()
	if err != nil {
		return err
	}

	gauge := *service.UniqueIPsMetric
	gauge.Set(float64(count))

	return nil
}

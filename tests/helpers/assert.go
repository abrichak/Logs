package helpers

import (
	"github.com/stretchr/testify/assert"
	"metrics-service/server"
	"metrics-service/server/services"
	"testing"
)

func AssertIPsCount(t *testing.T, s *server.Server, expectedIPsCount int) {
	t.Helper()

	count, _ := s.Redis.HLen(services.CacheKeyForIps).Result()
	assert.Equal(t, expectedIPsCount, int(count))
}

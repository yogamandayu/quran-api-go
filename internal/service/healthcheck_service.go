package service

import (
	"context"
	"quran-api-go/internal/domain/healthcheck"
	"time"
)

type healthCheckService struct{}

func NewHealthCheckService() healthcheck.HealthCheckService {
	return &healthCheckService{}
}

func (s *healthCheckService) HealthCheck(ctx context.Context) (healthcheck.HealthCheck, error) {
	return healthcheck.HealthCheck{
		Status:    "OK",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "-",
	}, nil
}

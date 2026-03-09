package service_test

import (
	"context"
	"quran-api-go/internal/service"
	"testing"
)

func TestHealthCheckService(t *testing.T) {
	service := service.NewHealthCheckService()
	health, err := service.HealthCheck(context.Background())
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if health.Status != "OK" {
		t.Fatalf("expected status OK, got %s", health.Status)
	}
}

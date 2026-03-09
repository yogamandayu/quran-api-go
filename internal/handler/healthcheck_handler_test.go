package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"quran-api-go/internal/domain/healthcheck"
	"quran-api-go/internal/handler"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// mockHealthCheckService is a test double for healthcheck.HealthCheckService.
type mockHealthCheckService struct {
	healthCheck func(ctx context.Context) (healthcheck.HealthCheck, error)
}

func (m *mockHealthCheckService) HealthCheck(ctx context.Context) (healthcheck.HealthCheck, error) {
	return m.healthCheck(ctx)
}

func newHealthCheckTestRouter(h *handler.HealthCheckHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", h.HealthCheck)
	return r
}

func TestHealthCheckHandler_List_OK(t *testing.T) {
	svc := &mockHealthCheckService{
		healthCheck: func(_ context.Context) (healthcheck.HealthCheck, error) {
			return healthcheck.HealthCheck{
				Status:    "OK",
				Timestamp: time.Now().UTC().Format(time.RFC3339),
				Version:   "v0.0.1",
			}, nil
		},
	}
	r := newHealthCheckTestRouter(handler.NewHealthCheckHandler(svc))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/health", nil))

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	status := body["status"]
	if status != "OK" {
		t.Fatalf("expected status field, got %v", body["status"])
	}
}

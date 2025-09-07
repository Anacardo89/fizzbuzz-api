package middleware

import (
	"net/http"
	"time"

	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
	"github.com/Anacardo89/fizzbuzz-api/pkg/obs"
)

type MiddlewareHandler struct {
	tokenManager *auth.TokenManager
	log          *logger.Logger
	writeTimeout time.Duration
	metrics      obs.MetricsClient
}

func NewMiddlewareHandler(tm *auth.TokenManager, l *logger.Logger, wto time.Duration, statsClient obs.MetricsClient) *MiddlewareHandler {
	return &MiddlewareHandler{
		tokenManager: tm,
		log:          l,
		writeTimeout: wto - time.Second,
		metrics:      statsClient,
	}
}

func (m *MiddlewareHandler) Wrap(next http.Handler) http.Handler {
	return m.Timeout(m.Tracing(m.Metrics(m.Log(next))))
}

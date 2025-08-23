package middleware

import (
	"net/http"
	"time"

	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

type MiddlewareHandler struct {
	tokenManager *auth.TokenManager
	log          *logger.Logger
	writeTimeout time.Duration
}

func NewMiddlewareHandler(tm *auth.TokenManager, l *logger.Logger, wto time.Duration) *MiddlewareHandler {
	return &MiddlewareHandler{
		tokenManager: tm,
		log:          l,
		writeTimeout: wto - time.Second,
	}
}

func (m *MiddlewareHandler) Wrap(next http.Handler) http.Handler {
	return m.Log(m.Timeout(next))
}

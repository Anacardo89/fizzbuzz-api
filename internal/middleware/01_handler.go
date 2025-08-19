package middleware

import (
	"github.com/Anacardo89/ecommerce_api/internal/auth"
	"github.com/Anacardo89/ecommerce_api/internal/logger"
)

type CtxKey string

const (
	UserIDKey CtxKey = "userID"
)

type MiddlewareHandler struct {
	tokenManager *auth.TokenManager
	logger       *logger.Logger
}

func NewMiddlewareHandler(tm *auth.TokenManager, l *logger.Logger) *MiddlewareHandler {
	return &MiddlewareHandler{
		tokenManager: tm,
		logger:       l,
	}
}

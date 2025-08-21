package middleware

import (
	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

type CtxKey string

const (
	UserIDKey CtxKey = "userID"
)

type MiddlewareHandler struct {
	tokenManager *auth.TokenManager
	log          *logger.Logger
}

func NewMiddlewareHandler(tm *auth.TokenManager, l *logger.Logger) *MiddlewareHandler {
	return &MiddlewareHandler{
		tokenManager: tm,
		log:          l,
	}
}

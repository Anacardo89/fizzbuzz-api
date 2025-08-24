package server

import (
	"log"
	"net/http/httptest"
	"time"

	"github.com/Anacardo89/fizzbuzz-api/config"
	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/internal/middleware"
	"github.com/Anacardo89/fizzbuzz-api/internal/repo"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

func LoadDefaultConfig() *config.Config {
	cfg := config.DefaultConfig()
	cfg.Server.ReadTimeout *= time.Second
	cfg.Server.WriteTimeout *= time.Second
	cfg.Server.ShutdownTimeout *= time.Second
	cfg.Token.Duration *= time.Minute
	cfg.DB.MaxConnLifetime *= time.Minute
	cfg.DB.MaxConnIdleTime *= time.Minute
	log.Printf("final cfg: %+v", cfg)
	return cfg
}

func NewMockServer() *httptest.Server {
	cfg := LoadDefaultConfig()
	l := logger.NewLogger(cfg.Log)
	tokenMan := auth.NewTokenManager(&cfg.Token)
	fbRepo := repo.NewMockFizzBuzzRepo()
	userRepo := repo.NewMockUserRepo()
	fh := api.NewFizzBuzzHandler(&cfg.Pag, fbRepo, l)
	ah := api.NewAuthHandler(tokenMan, userRepo, l)
	mw := middleware.NewMiddlewareHandler(tokenMan, l, cfg.Server.WriteTimeout)

	s := NewServer(&cfg.Server, l, fh, ah, mw)

	return httptest.NewServer(s.router)
}

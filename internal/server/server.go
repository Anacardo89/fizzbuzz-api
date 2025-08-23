package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Anacardo89/fizzbuzz-api/config"
	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/middleware"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

type Server struct {
	httpSrv  *http.Server
	router   http.Handler
	addr     string
	log      *logger.Logger
	timeouts ServerTimeouts
}

type ServerTimeouts struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func NewServer(cfg *config.ServerConfig, l *logger.Logger, fh *api.FizzBuzzHandler, ah *api.AuthHandler, mw *middleware.MiddlewareHandler) *Server {
	to := ServerTimeouts{
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		ShutdownTimeout: cfg.ShutdownTimeout,
	}
	s := &Server{
		router:   NewRouter(fh, ah, mw),
		addr:     fmt.Sprintf(":%s", cfg.Port),
		log:      l,
		timeouts: to,
	}
	return s
}

func (s *Server) Start() error {
	s.httpSrv = &http.Server{
		Addr:         s.addr,
		Handler:      s.router,
		ReadTimeout:  s.timeouts.ReadTimeout,
		WriteTimeout: s.timeouts.WriteTimeout,
	}
	s.log.Info("Starting server on", "adress", s.addr)
	return s.httpSrv.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeouts.ShutdownTimeout)
	defer cancel()
	if s.httpSrv != nil {
		return s.httpSrv.Shutdown(ctx)
	}
	return nil
}

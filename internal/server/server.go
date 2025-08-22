package server

import (
	"context"
	"net/http"
	"time"

	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/internal/middleware"
	"github.com/Anacardo89/fizzbuzz-api/internal/repo"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
	"github.com/gorilla/mux"
)

type Server struct {
	httpSrv  *http.Server
	router   *mux.Router
	addr     string
	log      *logger.Logger
	timeouts ServerTimeouts
}

type ServerTimeouts struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func NewServer(addr string, fbRepo *repo.FizzBuzzRepo, userRepo *repo.UserRepo, l *logger.Logger, to ServerTimeouts, tm *auth.TokenManager) *Server {
	fh := api.NewFizzBuzzHandler(fbRepo, l)
	ah := api.NewAuthHandler(tm, userRepo, l)
	mw := middleware.NewMiddlewareHandler(tm, l)
	s := &Server{
		router:   NewRouter(fh, ah, mw),
		addr:     addr,
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

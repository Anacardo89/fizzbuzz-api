package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	httpSrv  *http.Server
	Router   *mux.Router
	Addr     string
	DB       *sql.DB
	Logger   *log.Logger
	Timeouts ServerTimeouts
}

type ServerTimeouts struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func NewServer(addr string, db *sql.DB, logger *log.Logger, timeouts ServerTimeouts) *Server {
	s := &Server{
		Router:   mux.NewRouter(),
		Addr:     addr,
		DB:       db,
		Logger:   logger,
		Timeouts: timeouts,
	}
	return s
}

func (s *Server) Start() error {
	s.httpSrv = &http.Server{
		Addr:         s.Addr,
		Handler:      s.Router,
		ReadTimeout:  s.Timeouts.ReadTimeout,
		WriteTimeout: s.Timeouts.WriteTimeout,
	}

	s.Logger.Println("Starting server on", s.Addr)
	return s.httpSrv.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeouts.ShutdownTimeout)
	defer cancel()

	if s.DB != nil {
		s.DB.Close()
	}
	if s.httpSrv != nil {
		return s.httpSrv.Shutdown(ctx)
	}

	return nil
}

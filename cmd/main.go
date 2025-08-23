package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Anacardo89/fizzbuzz-api/config"
	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/internal/middleware"
	"github.com/Anacardo89/fizzbuzz-api/internal/server"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

func main() {
	// Dependencies
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	logg := logger.NewLogger(cfg.Log)
	logg.Info("config", "config", cfg)
	tokenMan := auth.NewTokenManager(&cfg.Token)
	fbRepo, userRepo, err := initDB(cfg.DB)
	if err != nil {
		logg.Fatal("failed to init db: %v", err)
	}
	defer fbRepo.Close()
	defer userRepo.Close()
	fh := api.NewFizzBuzzHandler(&cfg.Pag, fbRepo, logg)
	ah := api.NewAuthHandler(tokenMan, userRepo, logg)
	mw := middleware.NewMiddlewareHandler(tokenMan, logg, cfg.Server.WriteTimeout)

	srv := server.NewServer(&cfg.Server, logg, fh, ah, mw)

	// Serve
	stopChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		errChan <- srv.Start()
	}()
	select {
	case sig := <-stopChan:
		logg.Info("Shutting down...", "signal", sig)
		if err := srv.Shutdown(); err != nil {
			logg.Fatal("Failed to shutdown server gracefully", "error", err)
		}
		logg.Info("Server stopped gracefully")
	case err := <-errChan:
		if err != http.ErrServerClosed {
			logg.Fatal("Server failed", "error", err)
		}
	}
}

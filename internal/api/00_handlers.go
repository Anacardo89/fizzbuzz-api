package api

import (
	"github.com/Anacardo89/fizzbuzz-api/config"
	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/internal/repo"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

// FizzBuzz

type FizzBuzzHandler struct {
	cfg *config.PagConfig
	db  *repo.FizzBuzzRepo
	log *logger.Logger
}

func NewFizzBuzzHandler(cfg *config.PagConfig, r *repo.FizzBuzzRepo, l *logger.Logger) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		db:  r,
		log: l,
	}
}

// Auth

type AuthHandler struct {
	tokenManger *auth.TokenManager
	db          *repo.UserRepo
	log         *logger.Logger
}

func NewAuthHandler(tm *auth.TokenManager, r *repo.UserRepo, l *logger.Logger) *AuthHandler {
	return &AuthHandler{
		tokenManger: tm,
		db:          r,
		log:         l,
	}
}

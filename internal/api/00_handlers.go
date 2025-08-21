package api

import (
	"github.com/Anacardo89/fizzbuzz-api/internal/auth"
	"github.com/Anacardo89/fizzbuzz-api/internal/repo"
	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

// FizzBuzz

type FizzBuzzHandler struct {
	repo *repo.FizzBuzzRepo
	log  *logger.Logger
}

func NewFizzBuzzHandler(r *repo.FizzBuzzRepo, l *logger.Logger) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		repo: r,
		log:  l,
	}
}

// Auth

type AuthHandler struct {
	tokenManger *auth.TokenManager
	repo        *repo.UserRepo
	log         *logger.Logger
}

func NewAuthHandler(tm *auth.TokenManager, r *repo.UserRepo, l *logger.Logger) *AuthHandler {
	return &AuthHandler{
		tokenManger: tm,
		repo:        r,
		log:         l,
	}
}

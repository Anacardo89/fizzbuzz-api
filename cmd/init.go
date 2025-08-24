package main

import (
	"github.com/Anacardo89/fizzbuzz-api/config"
	"github.com/Anacardo89/fizzbuzz-api/internal/repo"
	"github.com/Anacardo89/fizzbuzz-api/pkg/db"
)

func initDB(cfg config.DBConfig) (repo.FizzBuzzRepo, repo.UserRepo, error) {
	fbPool, err := db.Connect(cfg)
	if err != nil {
		return nil, nil, err
	}
	userPool, err := db.Connect(cfg)
	if err != nil {
		return nil, nil, err
	}
	fbRepo := repo.NewFizzBuzzRepo(fbPool)
	userRepo := repo.NewUserRepo(userPool)
	return fbRepo, userRepo, nil
}

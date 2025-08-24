package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// FizzBuzz

type fizzBuzzHandler struct {
	pool *pgxpool.Pool
}

func NewFizzBuzzRepo(pool *pgxpool.Pool) FizzBuzzRepo {
	return &fizzBuzzHandler{
		pool: pool,
	}
}

func (r *fizzBuzzHandler) Close() {
	r.pool.Close()
}

// User

type userHandler struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) UserRepo {
	return &userHandler{
		pool: pool,
	}
}

func (r *userHandler) Close() {
	r.pool.Close()
}

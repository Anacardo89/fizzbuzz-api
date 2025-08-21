package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// FizzBuzz

type FizzBuzzRepo struct {
	pool *pgxpool.Pool
}

func NewFizzBuzzRepo(pool *pgxpool.Pool) *FizzBuzzRepo {
	return &FizzBuzzRepo{
		pool: pool,
	}
}

func (r *FizzBuzzRepo) Close() {
	r.pool.Close()
}

// User

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		pool: pool,
	}
}

func (r *UserRepo) Close() {
	r.pool.Close()
}

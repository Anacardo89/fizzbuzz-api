package api

import (
	"errors"

	"github.com/google/uuid"
)

// Error

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	ErrInvalidLimit      = errors.New("limit must be greater than 0 and reasonable (<= 1.000.000)")
	ErrIntLessThan1      = errors.New("int1 and int2 must be > 0")
	ErrEmptyString       = errors.New("str1 and str2 cannot be empty")
	ErrFieldNotInt       = "%s must be a valid integer"
	ErrInvalidPayload    = errors.New("invalid payload")
	ErrInvalidLoginCreds = errors.New("invalid username or password")
	ErrInternalError     = errors.New("internal error")
)

// FizzBuzz

type FizzBuzzResponse struct {
	Payload []string `json:"payload"`
}

// Stats

type FizzBuzzStatsResponse struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Count int    `json:"count"`
}

// Auth

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	UserID uuid.UUID `json:"user_id"`
}

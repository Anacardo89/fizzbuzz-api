package api

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Anacardo89/ecommerce_api/internal/repo"
)

var (
	ErrInvalidLimit = errors.New("limit must be greater than 0 and reasonable (<= 1.000.000)")
	ErrIntLessThan1 = errors.New("int1 and int2 must be > 0")
	ErrEmptyString  = errors.New("str1 and str2 cannot be empty")
	ErrFieldNotInt  = "%s must be a valid integer"
)

type FizzBuzzParams struct {
	Int1  int
	Int2  int
	Str1  string
	Str2  string
	Limit int
}

func NewFizzBuzzParams(int1Str, int2Str, str1, str2, limitStr string) (*FizzBuzzParams, error) {
	int1, err := strconv.Atoi(int1Str)
	if err != nil {
		return nil, fmt.Errorf(ErrFieldNotInt, "int1")
	}
	int2, err := strconv.Atoi(int2Str)
	if err != nil {
		return nil, fmt.Errorf(ErrFieldNotInt, "int2")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, fmt.Errorf(ErrFieldNotInt, "limit")
	}
	p := FizzBuzzParams{
		Int1:  int1,
		Int2:  int2,
		Str1:  str1,
		Str2:  str2,
		Limit: limit,
	}
	if err := ValidateFizzBuzzParams(p); err != nil {
		return nil, err
	}
	return &p, nil
}

func ValidateFizzBuzzParams(p FizzBuzzParams) error {
	if p.Limit <= 0 || p.Limit > 1_000_000 {
		return ErrInvalidLimit
	}
	if p.Int1 <= 0 || p.Int2 <= 0 {
		return ErrIntLessThan1
	}
	if p.Str1 == "" || p.Str2 == "" {
		return ErrEmptyString
	}
	return nil
}

func ParamsFromDB(row repo.FizzBuzzRow) (FizzBuzzParams, int) {
	return FizzBuzzParams{
		Int1:  row.Int1,
		Int2:  row.Int2,
		Str1:  row.Str1,
		Str2:  row.Str2,
		Limit: 0,
	}, row.RequestCount
}

func ParamsToDB(p FizzBuzzParams) repo.FizzBuzzRow {
	return repo.FizzBuzzRow{
		Int1: p.Int1,
		Int2: p.Int2,
		Str1: p.Str1,
		Str2: p.Str2,
	}
}

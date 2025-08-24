package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Anacardo89/fizzbuzz-api/internal/repo"
)

// FizzBuzz

type FizzBuzzURLParams struct {
	Int1  int
	Int2  int
	Str1  string
	Str2  string
	Limit int
}

func NewFizzBuzzParams(int1Str, int2Str, str1, str2, limitStr string) (*FizzBuzzURLParams, error) {
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
	p := FizzBuzzURLParams{
		Int1:  int1,
		Int2:  int2,
		Str1:  strings.ToLower(str1),
		Str2:  strings.ToLower(str2),
		Limit: limit,
	}
	if err := ValidateFizzBuzzParams(p); err != nil {
		return nil, err
	}
	return &p, nil
}

func ValidateFizzBuzzParams(p FizzBuzzURLParams) error {
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

func ParamsToDB(p FizzBuzzURLParams) repo.FizzBuzzRow {
	return repo.FizzBuzzRow{
		Int1: p.Int1,
		Int2: p.Int2,
		Str1: p.Str1,
		Str2: p.Str2,
	}
}

// AllStats

func (h *FizzBuzzHandler) ValidateAllStatsInput(offsetStr, limitStr string) (int, int, error) {
	offset := 0
	h.log.Info(fmt.Sprintf("offsetStr: %s", offsetStr))
	h.log.Info(fmt.Sprintf("limitStr: %s", limitStr))
	limit := h.cfg.DefaultPageSize
	maxLimit := h.cfg.MaxPageSize
	o, err := strconv.Atoi(offsetStr)
	if err != nil {
		return offset, limit, fmt.Errorf(ErrFieldNotInt, "offset")
	} else if o < 0 && offsetStr != "" {
		err = errors.New("invalid offset, must be >= 0")
		return offset, limit, err
	}
	offset = o
	l, err := strconv.Atoi(limitStr)
	if err != nil {
		return offset, limit, fmt.Errorf(ErrFieldNotInt, "limit")
	} else if l <= 0 && limitStr != "" {
		err = errors.New("invalid limit, must be > 0")
		return offset, limit, err
	}
	if l > maxLimit {
		limit = maxLimit
	} else {
		limit = l
	}
	return offset, limit, nil
}

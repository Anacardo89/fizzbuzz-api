package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

type FizzBuzzRow struct {
	ID           uuid.UUID
	Int1         int
	Int2         int
	Str1         string
	Str2         string
	RequestCount int
}

func (r *FizzBuzzRepo) UpsertFizzBuzz(ctx context.Context, params FizzBuzzRow) error {
	query := `
		INSERT INTO fizzbuzz (
			int1,
			int2,
			str1,
			str2
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (int1, int2, str1, str2)
		DO UPDATE SET request_count = fizzbuzz.request_count + 1
	;`
	if _, err := r.pool.Exec(ctx, query,
		params.Int1,
		params.Int2,
		params.Str1,
		params.Str2,
	); err != nil {
		return fmt.Errorf("failed to upsert fizzbuzz: %w", err)
	}
	return nil
}

// GetMostUsed returns the record with the highest request_count
func (r *FizzBuzzRepo) GetMostUsed(ctx context.Context) (*FizzBuzzRow, error) {
	query := `
		SELECT 
			int1,
			int2,
			str1,
			str2,
			request_count
		FROM fizzbuzz
		ORDER BY request_count DESC
		LIMIT 1
	;`
	var fbrow FizzBuzzRow
	row := r.pool.QueryRow(ctx, query)
	if err := row.Scan(
		&fbrow.Int1,
		&fbrow.Int2,
		&fbrow.Str1,
		&fbrow.Str2,
		&fbrow.RequestCount,
	); err != nil {
		return nil, fmt.Errorf("failed to get most used: %w", err)
	}
	return &fbrow, nil
}

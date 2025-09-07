package repo

import (
	"context"
	"fmt"

	"github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
	"github.com/google/uuid"
)

type FizzBuzzRepo interface {
	Close()
	UpsertFizzBuzz(ctx context.Context, params FizzBuzzRow) error
	SelectTopFizzBuzzQuery(ctx context.Context) (*FizzBuzzRow, error)
	SelectFizzBuzzQueries(ctx context.Context, limit, offset int) ([]FizzBuzzRow, error)
}

type FizzBuzzRow struct {
	ID           uuid.UUID
	Int1         int
	Int2         int
	Str1         string
	Str2         string
	RequestCount int
}

func (r *fizzBuzzHandler) UpsertFizzBuzz(ctx context.Context, params FizzBuzzRow) error {
	span, spanCtx := tracer.StartSpanFromContext(ctx, "db.upsert_fizzbuzz")
	defer span.Finish()

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
	if _, err := r.pool.Exec(spanCtx, query,
		params.Int1,
		params.Int2,
		params.Str1,
		params.Str2,
	); err != nil {
		return fmt.Errorf("failed to upsert fizzbuzz: %w", err)
	}
	return nil
}

func (r *fizzBuzzHandler) SelectTopFizzBuzzQuery(ctx context.Context) (*FizzBuzzRow, error) {
	span, spanCtx := tracer.StartSpanFromContext(ctx, "db.select_top_fizzbuzz")
	defer span.Finish()

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
	row := r.pool.QueryRow(spanCtx, query)
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

func (r *fizzBuzzHandler) SelectFizzBuzzQueries(ctx context.Context, limit, offset int) ([]FizzBuzzRow, error) {
	span, spanCtx := tracer.StartSpanFromContext(ctx, "db.select_fizzbuzz")
	defer span.Finish()

	query := `
		SELECT 
			int1,
			int2,
			str1,
			str2,
			request_count
		FROM fizzbuzz
		ORDER BY request_count DESC
		LIMIT $1
		OFFSET $2;
	;`
	rows, err := r.pool.Query(spanCtx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query fizzbuzz: %w", err)
	}
	defer rows.Close()
	var results []FizzBuzzRow
	for rows.Next() {
		var fbrow FizzBuzzRow
		if err := rows.Scan(
			&fbrow.Int1,
			&fbrow.Int2,
			&fbrow.Str1,
			&fbrow.Str2,
			&fbrow.RequestCount,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, fbrow)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}
	return results, nil
}

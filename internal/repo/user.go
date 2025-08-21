package repo

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

type UserRow struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (r *UserRepo) InsertUser(ctx context.Context, username, hashedPassword string) (uuid.UUID, error) {
	const query = `
		INSERT INTO users (
			username,
			password
		)
		VALUES ($1, $2)
		RETURNING id;
	;`
	var ID uuid.UUID
	if err := r.pool.QueryRow(ctx, query,
		username,
		hashedPassword,
	).Scan(&ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // unique_violation
			return uuid.Nil, ErrUserExists
		}
		return uuid.Nil, err
	}
	return ID, nil
}

func (r *UserRepo) SelectUser(ctx context.Context, username string) (*UserRow, error) {
	const query = `
		SELECT 
			id,
			username,
			password
		FROM users
		WHERE username = $1
	;`
	user := UserRow{}
	if err := r.pool.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	); err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

package repo

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type MockUserRepo struct {
	mu    sync.Mutex
	users map[string]*UserRow // key = username
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{
		users: make(map[string]*UserRow),
	}
}

func (r *MockUserRepo) Close() {}

func (r *MockUserRepo) InsertUser(ctx context.Context, username, hashedPassword string) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[username]; exists {
		return uuid.Nil, ErrUserExists
	}
	ID := uuid.New()
	r.users[username] = &UserRow{
		ID:       ID.String(),
		Username: username,
		Password: hashedPassword,
	}
	return ID, nil
}

func (r *MockUserRepo) SelectUser(ctx context.Context, username string) (*UserRow, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[username]
	if !exists {
		return nil, ErrUserNotFound
	}
	return &UserRow{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

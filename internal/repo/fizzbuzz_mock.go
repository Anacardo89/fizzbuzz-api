package repo

import (
	"context"
	"fmt"
	"sort"
	"sync"
)

type MockFizzBuzzRepo struct {
	mu   sync.Mutex
	rows map[string]FizzBuzzRow // key = int1:int2:str1:str2
}

func NewMockFizzBuzzRepo() *MockFizzBuzzRepo {
	return &MockFizzBuzzRepo{
		rows: make(map[string]FizzBuzzRow),
	}
}

func key(r FizzBuzzRow) string {
	return fmt.Sprintf("%d:%d:%s:%s", r.Int1, r.Int2, r.Str1, r.Str2)
}

func (m *MockFizzBuzzRepo) Close() {}

func (m *MockFizzBuzzRepo) UpsertFizzBuzz(ctx context.Context, row FizzBuzzRow) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	k := key(row)
	if existing, ok := m.rows[k]; ok {
		existing.RequestCount++
		m.rows[k] = existing
	} else {
		row.RequestCount = 1
		m.rows[k] = row
	}
	return nil
}

func (m *MockFizzBuzzRepo) SelectTopFizzBuzzQuery(ctx context.Context) (*FizzBuzzRow, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.rows) == 0 {
		return nil, nil
	}
	var top FizzBuzzRow
	maxCount := 0
	for _, r := range m.rows {
		if r.RequestCount > maxCount {
			maxCount = r.RequestCount
			top = r
		}
	}
	return &top, nil
}

func (m *MockFizzBuzzRepo) SelectFizzBuzzQueries(ctx context.Context, limit, offset int) ([]FizzBuzzRow, error) {
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 0
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	rows := make([]FizzBuzzRow, 0, len(m.rows))
	for _, r := range m.rows {
		rows = append(rows, r)
	}
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].RequestCount > rows[j].RequestCount
	})
	if offset >= len(rows) {
		return []FizzBuzzRow{}, nil
	}
	end := offset + limit
	if end > len(rows) {
		end = len(rows)
	}
	return rows[offset:end], nil
}

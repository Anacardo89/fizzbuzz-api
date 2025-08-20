package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name   string
		int1   int
		int2   int
		str1   string
		str2   string
		limit  int
		expect []string
	}{
		{
			name:   "fizzbuzz",
			int1:   3,
			int2:   5,
			str1:   "fizz",
			str2:   "buzz",
			limit:  16,
			expect: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"},
		},
		{
			name:   "different divisors and strings",
			int1:   2,
			int2:   4,
			str1:   "foo",
			str2:   "bar",
			limit:  8,
			expect: []string{"1", "foo", "3", "foobar", "5", "foo", "7", "foobar"},
		},
		{
			name:   "limit = 1",
			int1:   3,
			int2:   5,
			limit:  1,
			str1:   "fizz",
			str2:   "buzz",
			expect: []string{"1"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := FizzBuzz(tc.int1, tc.int2, tc.str1, tc.str2, tc.limit)
			assert.Equal(t, tc.expect, result)
		})
	}
}

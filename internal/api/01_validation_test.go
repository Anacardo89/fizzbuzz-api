package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateFizzBuzzParams(t *testing.T) {
	tests := []struct {
		name    string
		params  FizzBuzzURLParams
		wantErr error
	}{
		{
			name: "valid params",
			params: FizzBuzzURLParams{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			wantErr: nil,
		},
		{
			name: "invalid limit zero",
			params: FizzBuzzURLParams{
				Int1:  3,
				Int2:  5,
				Limit: 0,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			wantErr: ErrInvalidLimit,
		},
		{
			name: "invalid limit too high",
			params: FizzBuzzURLParams{
				Int1:  3,
				Int2:  5,
				Limit: 1_000_001,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			wantErr: ErrInvalidLimit,
		},
		{
			name: "invalid Int1",
			params: FizzBuzzURLParams{
				Int1:  -1,
				Int2:  5,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			wantErr: ErrIntLessThan1,
		},
		{
			name: "invalid Int2",
			params: FizzBuzzURLParams{
				Int1:  3,
				Int2:  0,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			wantErr: ErrIntLessThan1,
		},
		{
			name: "invalid ints",
			params: FizzBuzzURLParams{
				Int1:  -1,
				Int2:  0,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			wantErr: ErrIntLessThan1,
		},
		{
			name: "empty Str1",
			params: FizzBuzzURLParams{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "",
				Str2:  "buzz",
			},
			wantErr: ErrEmptyString,
		},
		{
			name: "empty Str2",
			params: FizzBuzzURLParams{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "",
			},
			wantErr: ErrEmptyString,
		},
		{
			name: "empty strings",
			params: FizzBuzzURLParams{
				Int1:  3,
				Int2:  5,
				Limit: 100,
				Str1:  "",
				Str2:  "",
			},
			wantErr: ErrEmptyString,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFizzBuzzParams(tt.params)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestNewFizzBuzzParams(t *testing.T) {
	tests := []struct {
		name      string
		int1      string
		int2      string
		str1      string
		str2      string
		limit     string
		wantErr   bool
		errString string
	}{
		{
			name:  "valid params",
			int1:  "3",
			int2:  "5",
			str1:  "fizz",
			str2:  "buzz",
			limit: "100",
		},
		{
			name:      "invalid int1 not an int",
			int1:      "abc",
			int2:      "5",
			str1:      "fizz",
			str2:      "buzz",
			limit:     "100",
			wantErr:   true,
			errString: fmt.Sprintf(ErrFieldNotInt, "int1"),
		},
		{
			name:      "invalid int2 not an int",
			int1:      "3",
			int2:      "foo",
			str1:      "fizz",
			str2:      "buzz",
			limit:     "100",
			wantErr:   true,
			errString: fmt.Sprintf(ErrFieldNotInt, "int2"),
		},
		{
			name:      "invalid limit not an int",
			int1:      "3",
			int2:      "5",
			str1:      "fizz",
			str2:      "buzz",
			limit:     "bar",
			wantErr:   true,
			errString: fmt.Sprintf(ErrFieldNotInt, "limit"),
		},
		{
			name:      "limit too big",
			int1:      "3",
			int2:      "5",
			str1:      "fizz",
			str2:      "buzz",
			limit:     "1000001",
			wantErr:   true,
			errString: ErrInvalidLimit.Error(),
		},
		{
			name:      "ints not positive",
			int1:      "0",
			int2:      "-5",
			str1:      "fizz",
			str2:      "buzz",
			limit:     "100",
			wantErr:   true,
			errString: ErrIntLessThan1.Error(),
		},
		{
			name:      "empty strings",
			int1:      "3",
			int2:      "5",
			str1:      "",
			str2:      "",
			limit:     "100",
			wantErr:   true,
			errString: ErrEmptyString.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := NewFizzBuzzParams(tt.int1, tt.int2, tt.str1, tt.str2, tt.limit)

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errString)
				require.Nil(t, params)
			} else {
				require.NoError(t, err)
				require.NotNil(t, params)
				require.Equal(t, tt.int1, fmt.Sprint(params.Int1))
				require.Equal(t, tt.int2, fmt.Sprint(params.Int2))
				require.Equal(t, tt.str1, params.Str1)
				require.Equal(t, tt.str2, params.Str2)
				require.Equal(t, tt.limit, fmt.Sprint(params.Limit))
			}
		})
	}
}

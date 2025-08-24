package integrationtest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/server"
	"github.com/stretchr/testify/require"
)

func TestGetFizzBuzz(t *testing.T) {
	srv := server.NewMockServer()
	defer srv.Close()
	url := fmt.Sprintf("%s/fizzbuzz?int1=3&int2=5&str1=Fizz&str2=Buzz&limit=16", srv.URL)
	resp, err := http.Get(url)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var got api.FizzBuzzResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	require.NoError(t, err)
	require.Len(t, got.Payload, 16)
	require.Equal(t, "Fizz", got.Payload[2])
	require.Equal(t, "Buzz", got.Payload[4])
	require.Equal(t, "FizzBuzz", got.Payload[14])
}

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

func TestHealthCheck(t *testing.T) {
	srv := server.NewMockServer()
	defer srv.Close()
	resp, err := http.Get(srv.URL)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var got api.HealthCheckResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	require.NoError(t, err)
	require.Equal(t, "OK", got.Status)
}

func TestCatchAll(t *testing.T) {
	srv := server.NewMockServer()
	defer srv.Close()
	resp, err := http.Get(fmt.Sprintf("%s/does-not-exist", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	var got api.ErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	require.NoError(t, err)
	require.Equal(t, "endpoint not found", got.Error)
}

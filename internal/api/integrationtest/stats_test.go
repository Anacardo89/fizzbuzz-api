package integrationtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/server"
	"github.com/stretchr/testify/require"
)

func TestGetStatsTopQuery(t *testing.T) {
	srv := server.NewMockServer()
	defer srv.Close()
	url := fmt.Sprintf("%s/fizzbuzz?int1=3&int2=5&str1=Fizz&str2=Buzz&limit=16", srv.URL)
	http.Get(url)
	url = fmt.Sprintf("%s/stats", srv.URL)
	resp, err := http.Get(url)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var got api.StatsResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	require.NoError(t, err)
	require.Equal(t, 3, got.Int1)
	require.Equal(t, 5, got.Int2)
	require.Equal(t, "Fizz", got.Str1)
	require.Equal(t, "Buzz", got.Str2)
	require.Equal(t, 1, got.Hits)
}

func TestGetStatsAllQueries(t *testing.T) {
	srv := server.NewMockServer()
	defer srv.Close()
	username := "charlie"
	password := "mypassword"
	reqBody := api.AuthRequest{
		Username: username,
		Password: password,
	}
	url := fmt.Sprintf("%s/fizzbuzz?int1=3&int2=5&str1=Fizz&str2=Buzz&limit=16", srv.URL)
	http.Get(url)
	url = fmt.Sprintf("%s/fizzbuzz?int1=3&int2=5&str1=Fizz&str2=Buzz&limit=16", srv.URL)
	http.Get(url)
	url = fmt.Sprintf("%s/fizzbuzz?int1=2&int2=4&str1=Foo&str2=Bar&limit=16", srv.URL)
	http.Get(url)
	// Register
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(reqBody)
	url = fmt.Sprintf("%s/auth/register", srv.URL)
	resp, err := http.Post(url, "application/json", buf)
	require.NoError(t, err)
	resp.Body.Close()
	// Login
	buf = new(bytes.Buffer)
	json.NewEncoder(buf).Encode(reqBody)
	resp, err = http.Post(fmt.Sprintf("%s/auth/login", srv.URL), "application/json", buf)
	require.NoError(t, err)
	body := api.LoginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	require.NoError(t, err)
	resp.Body.Close()
	// Stats Request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/stats/all?offset=0&limit=5", srv.URL), nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", body.Token))
	client := &http.Client{}
	resp, err = client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var got api.AllStatsResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	require.NoError(t, err)
	require.LessOrEqual(t, len(got.Stats), 5)
	require.Equal(t, 0, got.Offset)
	require.Equal(t, 5, got.Limit)
	require.Equal(t, len(got.Stats), got.StatsLen)
	if len(got.Stats) > 0 {
		require.Equal(t, 3, got.Stats[0].Int1)
		require.Equal(t, "Fizz", got.Stats[0].Str1)
	}
}

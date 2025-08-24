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

func TestAuthEndpoints(t *testing.T) {
	srv := server.NewMockServer()
	defer srv.Close()

	t.Run("Register Success", func(t *testing.T) {
		reqBody := api.AuthRequest{
			Username: "alice",
			Password: "password123",
		}
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(reqBody)
		url := fmt.Sprintf("%s/auth/register", srv.URL)
		resp, err := http.Post(url, "application/json", buf)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusCreated, resp.StatusCode)
		var got api.RegisterResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		require.NoError(t, err)
		require.NotEmpty(t, got.UserID)
	})

	t.Run("Register Conflict", func(t *testing.T) {
		username := "bob"
		password := "pass123"
		reqBody := api.AuthRequest{
			Username: username,
			Password: password,
		}
		// Register
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(reqBody)
		url := fmt.Sprintf("%s/auth/register", srv.URL)
		resp, err := http.Post(url, "application/json", buf)
		require.NoError(t, err)
		resp.Body.Close()
		// Register twice
		buf = new(bytes.Buffer)
		json.NewEncoder(buf).Encode(reqBody)
		resp, err = http.Post(fmt.Sprintf("%s/auth/register", srv.URL), "application/json", buf)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusConflict, resp.StatusCode)
		var got api.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		require.NoError(t, err)
		require.Contains(t, got.Error, "username already exists")
	})

	t.Run("Login Success", func(t *testing.T) {
		username := "charlie"
		password := "mypassword"
		reqBody := api.AuthRequest{
			Username: username,
			Password: password,
		}
		// Register
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(reqBody)
		url := fmt.Sprintf("%s/auth/register", srv.URL)
		resp, err := http.Post(url, "application/json", buf)
		require.NoError(t, err)
		resp.Body.Close()
		// Login
		buf = new(bytes.Buffer)
		json.NewEncoder(buf).Encode(reqBody)
		url = fmt.Sprintf("%s/auth/login", srv.URL)
		resp, err = http.Post(url, "application/json", buf)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)
		var got api.LoginResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		require.NoError(t, err)
		require.NotEmpty(t, got.Token)
	})

	t.Run("Login Wrong Password", func(t *testing.T) {
		username := "dave"
		password := "securepass"
		reqBody := api.AuthRequest{
			Username: username,
			Password: password,
		}
		// Register
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(reqBody)
		url := fmt.Sprintf("%s/auth/register", srv.URL)
		resp, err := http.Post(url, "application/json", buf)
		require.NoError(t, err)
		resp.Body.Close()
		// Login
		reqBody = api.AuthRequest{
			Username: username,
			Password: "wrongpass",
		}
		buf = new(bytes.Buffer)
		json.NewEncoder(buf).Encode(reqBody)
		url = fmt.Sprintf("%s/auth/login", srv.URL)
		resp, err = http.Post(url, "application/json", buf)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		var got api.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		require.NoError(t, err)
		require.Contains(t, got.Error, "invalid username or password")
	})
}

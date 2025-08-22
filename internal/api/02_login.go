package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Anacardo89/fizzbuzz-api/pkg/crypto"
)

type FailHelper struct {
	LogMsg string
	Err    error
	Status int
	OutMsg string
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Error Handling
	fail := func(logMsg string, e error, status int, outMsg string) {
		h.log.Error(logMsg, "error", e,
			"status_code", status,
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"client_ip", r.RemoteAddr,
		)
		w.WriteHeader(status)
		resp := ErrorResponse{Error: outMsg}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			h.log.Error("failed to encode error response", "error", err)
		}
	}
	//

	// Execution
	w.Header().Set("Content-Type", "application/json")
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fail("invalid login payload", err, http.StatusBadRequest, ErrInvalidPayload.Error())
		return
	}
	user, err := h.repo.SelectUser(r.Context(), req.Username)
	if err != nil {
		fail("dberr: failed to get user", err, http.StatusUnauthorized, ErrInvalidLoginCreds.Error())
		return
	}
	if !crypto.ValidatePassword(user.Password, req.Password) {
		fail("invalid password", errors.New("password and hash do not match"), http.StatusUnauthorized, ErrInvalidLoginCreds.Error())
		return
	}
	token, err := h.tokenManger.GenerateToken(user.ID)
	if err != nil {
		fail("failed to generate token", err, http.StatusInternalServerError, ErrInternalError.Error())
		return
	}
	resp := LoginResponse{
		Token: token,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fail("failed to encode response", err, http.StatusInternalServerError, ErrInternalError.Error())
	}
}

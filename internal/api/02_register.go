package api

import (
	"encoding/json"
	"net/http"

	"github.com/Anacardo89/fizzbuzz-api/internal/repo"
	"github.com/Anacardo89/fizzbuzz-api/pkg/crypto"
)

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
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
		fail("invalid register payload", err, http.StatusBadRequest, ErrInvalidPayload.Error())
		return
	}
	hash, err := crypto.HashPassword(req.Password)
	if err != nil {
		fail("failed to hash password", err, http.StatusInternalServerError, ErrInternalError.Error())
		return
	}
	userID, err := h.db.InsertUser(r.Context(), req.Username, hash)
	if err != nil {
		if err == repo.ErrUserExists {
			fail("dberr: user exists", err, http.StatusConflict, "username already exists")
			return
		}
		fail("dberr: failed to insert user", err, http.StatusInternalServerError, ErrInternalError.Error())
		return
	}
	resp := RegisterResponse{
		UserID: userID,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fail("failed to encode response", err, http.StatusInternalServerError, ErrInternalError.Error())
	}
}

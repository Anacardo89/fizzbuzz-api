package api

import (
	"encoding/json"
	"net/http"

	"github.com/Anacardo89/fizzbuzz-api/internal/core"
)

type FizzBuzzResponse struct {
	Payload []string `json:"payload"`
}

func (h *FizzBuzzHandler) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	// Error Handling
	fail := func(msg string, e error, writeError bool, status int) {
		h.log.Error(msg, "error", e,
			"status_code", status,
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"client_ip", r.RemoteAddr,
		)
		if writeError {
			http.Error(w, e.Error(), status)
		}
	}
	//

	// Execution
	params, err := NewFizzBuzzParams(
		r.URL.Query().Get("int1"),
		r.URL.Query().Get("int2"),
		r.URL.Query().Get("str1"),
		r.URL.Query().Get("str2"),
		r.URL.Query().Get("limit"),
	)
	if err != nil {
		fail("invalid params", err, true, http.StatusBadRequest)
		return
	}

	result := core.FizzBuzz(params.Int1, params.Int2, params.Str1, params.Str2, params.Limit)

	paramsDB := ParamsToDB(*params)
	if err := h.repo.UpsertFizzBuzz(r.Context(), paramsDB); err != nil {
		fail("dberr: upsert fizzbuzz", err, false, 0)
	}

	resp := FizzBuzzResponse{
		Payload: result,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fail("failed to encode response", err, true, http.StatusInternalServerError)
	}
}

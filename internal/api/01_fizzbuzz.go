package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Anacardo89/fizzbuzz-api/internal/core"
)

func (h *FizzBuzzHandler) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	// Error Handling
	fail := func(logMsg string, e error, writeError bool, status int, outMsg string) {
		h.log.Error(logMsg, "error", e,
			"status_code", status,
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"client_ip", r.RemoteAddr,
		)
		if writeError {
			w.WriteHeader(status)
			resp := ErrorResponse{Error: outMsg}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				h.log.Error("failed to encode error response", "error", err)
			}
		}
	}
	//

	// Execution
	w.Header().Set("Content-Type", "application/json")
	params, err := NewFizzBuzzParams(
		r.URL.Query().Get("int1"),
		r.URL.Query().Get("int2"),
		r.URL.Query().Get("str1"),
		r.URL.Query().Get("str2"),
		r.URL.Query().Get("limit"),
	)
	if err != nil {
		fail("invalid params", err, true, http.StatusBadRequest, err.Error())
		return
	}
	result := core.FizzBuzz(params.Int1, params.Int2, params.Str1, params.Str2, params.Limit)
	resp := FizzBuzzResponse{
		Payload: result,
	}
	go func() {
		paramsDB := ParamsToDB(*params)
		if err := h.db.UpsertFizzBuzz(context.Background(), paramsDB); err != nil {
			fail("dberr: upsert fizzbuzz", err, false, 0, "")
		}
	}()
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fail("failed to encode response", err, true, http.StatusInternalServerError, ErrInternalError.Error())
	}
}

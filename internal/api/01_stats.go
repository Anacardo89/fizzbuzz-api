package api

import (
	"encoding/json"
	"net/http"
)

func (h *FizzBuzzHandler) GetTopQuery(w http.ResponseWriter, r *http.Request) {
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
	fb, err := h.repo.SelectTopFizzBuzzQuery(r.Context())
	if err != nil {
		fail("dberr: failed to fetch stats", err, http.StatusInternalServerError, ErrInternalError.Error())
		return
	}
	resp := FizzBuzzStatsResponse{
		Int1:  fb.Int1,
		Int2:  fb.Int2,
		Str1:  fb.Str1,
		Str2:  fb.Str2,
		Count: fb.RequestCount,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fail("failed to encode response", err, http.StatusInternalServerError, ErrInternalError.Error())
	}
}

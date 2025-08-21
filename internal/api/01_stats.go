package api

import (
	"encoding/json"
	"net/http"
)

type FizzBuzzStatsResponse struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Count int    `json:"count"`
}

func (h *FizzBuzzHandler) GetTopQuery(w http.ResponseWriter, r *http.Request) {
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
	fb, err := h.repo.SelectTopFizzBuzzQuery(r.Context())
	if err != nil {
		fail("dberr: failed to fetch stats", err, true, http.StatusInternalServerError)
		return
	}

	resp := FizzBuzzStatsResponse{
		Int1:  fb.Int1,
		Int2:  fb.Int2,
		Str1:  fb.Str1,
		Str2:  fb.Str2,
		Count: fb.RequestCount,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fail("failed to encode response", err, true, http.StatusInternalServerError)
	}
}

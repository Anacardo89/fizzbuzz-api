package api

import (
	"encoding/json"
	"net/http"
)

func (h *FizzBuzzHandler) GetStatsTopQuery(w http.ResponseWriter, r *http.Request) {
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
			h.log.Error("failed to encode error response body", "error", err)
		}
	}
	//

	// Execution
	w.Header().Set("Content-Type", "application/json")
	fb, err := h.db.SelectTopFizzBuzzQuery(r.Context())
	if err != nil {
		fail("dberr: failed to fetch stats", err, http.StatusInternalServerError, ErrInternalError.Error())
		return
	}
	body := StatsResponse{
		Int1: fb.Int1,
		Int2: fb.Int2,
		Str1: fb.Str1,
		Str2: fb.Str2,
		Hits: fb.RequestCount,
	}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fail("failed to encode response body", err, http.StatusInternalServerError, ErrInternalError.Error())
	}
}

func (h *FizzBuzzHandler) GetStatsAllQueries(w http.ResponseWriter, r *http.Request) {
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
			h.log.Error("failed to encode error response body", "error", err)
		}
	}
	//

	//Execution
	w.Header().Set("Content-Type", "application/json")
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")
	offset, limit, err := h.ValidateAllStatsInput(offsetStr, limitStr)
	if err != nil {
		fail("invalid params", err, http.StatusBadRequest, err.Error())
		return
	}
	statsDB, err := h.db.SelectFizzBuzzQueries(r.Context(), limit, offset)
	if err != nil {
		fail("dberr: failed to get stats", err, http.StatusInternalServerError, ErrInternalError.Error())
		return
	}
	stats := make([]StatsResponse, 0, len(statsDB))
	for _, s := range statsDB {
		stat := StatsResponse{
			Int1: s.Int1,
			Int2: s.Int2,
			Str1: s.Str1,
			Str2: s.Str2,
			Hits: s.RequestCount,
		}
		stats = append(stats, stat)
	}
	body := AllStatsResponse{
		Stats:    stats,
		StatsLen: len(stats),
		Offset:   offset,
		Limit:    limit,
	}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fail("failed to encode response body", err, http.StatusInternalServerError, ErrInternalError.Error())
	}
}

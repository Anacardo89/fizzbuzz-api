package middleware

import (
	"net/http"
	"time"
)

func (m *MiddlewareHandler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.log.Info("request received",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"time_received", start,
			"client_ip", r.RemoteAddr,
		)

		rw := newLogRW(w)
		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		m.log.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"status", rw.Status(),
			"size", rw.Size(),
			"duration_ms", duration.Milliseconds(),
			"client_ip", r.RemoteAddr,
		)
	})
}

// Used to capture status code of response
type LogRW struct {
	http.ResponseWriter
	status int
	size   int
}

func newLogRW(w http.ResponseWriter) *LogRW {
	return &LogRW{ResponseWriter: w}
}

func (rw *LogRW) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *LogRW) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

func (rw *LogRW) Status() int { return rw.status }
func (rw *LogRW) Size() int   { return rw.size }

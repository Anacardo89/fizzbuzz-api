package middleware

import (
	"net/http"
	"time"

	"github.com/Anacardo89/fizzbuzz-api/pkg/logger"
)

func (m *MiddlewareHandler) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		attrs := []any{
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"time_received", start,
			"client_ip", r.RemoteAddr,
		}
		attrs = append(attrs, logger.TraceAttrs(r.Context())...)
		m.log.Info("request received", attrs...)

		rw := getRWWrapper(w)
		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		attrs = append(attrs,
			"status", rw.Status(),
			"size", rw.Size(),
			"duration_ms", duration.Milliseconds(),
		)
		m.log.Info("request completed", attrs...)
	})
}

package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
)

func (m *MiddlewareHandler) Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ctx := tracer.StartSpanFromContext(r.Context(), r.URL.Path)
		defer span.Finish()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *MiddlewareHandler) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := getRWWrapper(w)
		next.ServeHTTP(rw, r)

		tags := []string{
			fmt.Sprintf("endpoint:%s", r.URL.Path),
			fmt.Sprintf("method:%s", r.Method),
			fmt.Sprintf("status:%d", rw.Status()),
		}

		m.metrics.Incr("http.requests_total", tags, 1)
		m.metrics.Timing("http.request_duration_ms", time.Since(start), tags, 1)
	})
}

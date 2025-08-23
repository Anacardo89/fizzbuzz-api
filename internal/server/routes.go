package server

import (
	"net/http"

	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(fh *api.FizzBuzzHandler, ah *api.AuthHandler, mw *middleware.MiddlewareHandler) http.Handler {
	r := mux.NewRouter()

	// Health check
	r.Handle("/", http.HandlerFunc(api.HealthCheck)).Methods("GET")

	// Auth
	r.Handle("/auth/register", http.HandlerFunc(ah.Register)).Methods("POST")
	r.Handle("/auth/login", http.HandlerFunc(ah.Login)).Methods("POST")

	// FizzBuzz
	r.Handle("/fizzbuzz", http.HandlerFunc(fh.GetFizzBuzz)).Methods("GET")

	// Stats
	r.Handle("/stats", http.HandlerFunc(fh.GetTopQuery)).Methods("GET")
	r.Handle("/stats/all", mw.Auth(http.HandlerFunc(fh.GetAllQueries))).Methods("GET")

	// Catch-all 404
	r.NotFoundHandler = http.HandlerFunc(api.CatchAll)

	return mw.Wrap(r)
}

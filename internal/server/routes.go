package server

import (
	"net/http"

	"github.com/Anacardo89/fizzbuzz-api/internal/api"
	"github.com/Anacardo89/fizzbuzz-api/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(fh *api.FizzBuzzHandler, ah *api.AuthHandler, mw *middleware.MiddlewareHandler) *mux.Router {
	r := mux.NewRouter()

	// Auth
	r.Handle("/auth/register", mw.Log(http.HandlerFunc(ah.Register))).Methods("POST")
	r.Handle("/auth/login", mw.Log(http.HandlerFunc(ah.Login))).Methods("POST")

	// FizzBuzz
	r.Handle("/fizzbuzz", mw.Log(http.HandlerFunc(fh.GetFizzBuzz))).Methods("GET")

	// Stats
	r.Handle("/fizzbuzz/stats", mw.Log(mw.Auth(http.HandlerFunc(fh.GetTopQuery)))).Methods("GET")
	r.Handle("/fizzbuzz/stats/all", mw.Log(mw.Auth(http.HandlerFunc(fh.GetAllQueries)))).Methods("GET")

	return r
}

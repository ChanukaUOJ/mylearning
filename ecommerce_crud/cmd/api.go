package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	// logger
	// db driver
}

// mount
func (app *application) mount() http.Handler {
	// gorilla mux / chi / fiber (server packages we can use)
	// here we are using 'chi'
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // important for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi."))
	})

	// http.ListenAndServe(":3333", r)
	return r
}

// run

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

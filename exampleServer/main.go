package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// setup of router
	r := chi.NewRouter()

	// basic middleware, good start
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// basic get call TODO: move out into it's own directory
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// this port WILL be used in prod
	http.ListenAndServe(":42069", r)
}

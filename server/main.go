package main

import (
	"net/http"
	"tempbin/server/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialization
	r := chi.NewRouter()
	go cleaner()

	// middlewares
	r.Use(middleware.Logger)

	// routes
	r.Post("/upload", handlers.Upload)

	http.ListenAndServe(":3000", r)
}

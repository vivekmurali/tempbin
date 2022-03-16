package main

import (
	"net/http"
	"tempbin/server/db"
	"tempbin/server/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialization
	r := chi.NewRouter()
	db.InitDB()
	// go cleaner()

	// middlewares
	r.Use(middleware.Logger)

	// routes
	r.Post("/upload", handlers.Upload)
	r.Route("/download", func(r chi.Router) {
		r.Get("/{url}", handlers.DownloadHandler)
		r.Post("/{url}", handlers.Download)
	})
	http.ListenAndServe(":3000", r)
}

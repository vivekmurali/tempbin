package main

import (
	"fmt"
	"net/http"
	"tempbin/server/db"
	"tempbin/server/handlers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-co-op/gocron"
)

func main() {
	// Initialization
	r := chi.NewRouter()
	db.InitDB()

	// go cleaner()
	s := gocron.NewScheduler(time.UTC)
	var job func() = worker
	s.Every(10).Minutes().Do(job)
	s.StartAsync()

	// middlewares
	r.Use(middleware.Logger)

	// routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})
	r.Get("/upload", handlers.UploadHandler)
	r.Post("/upload", handlers.Upload)
	r.Route("/download", func(r chi.Router) {
		r.Get("/{url}", handlers.DownloadHandler)
		r.Post("/{url}", handlers.Download)
	})
	http.ListenAndServe(":3000", r)
}

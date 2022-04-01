package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"tempbin/server/db"
	"tempbin/server/handlers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-co-op/gocron"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	prometheus.Register(handlers.TotalRequests)
	prometheus.Register(handlers.ResponseStatus)
	prometheus.Register(handlers.HttpDuration)
}

func main() {
	// Initialization
	r := chi.NewRouter()
	db.InitDB()

	// go cleaner()
	s := gocron.NewScheduler(time.UTC)
	var job func() = worker
	s.Every(1).Minutes().Do(job)
	s.StartAsync()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(handlers.PrometheusMiddleware)

	//Fileserver to serve static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "server/static"))
	FileServer(r, "/static", filesDir)

	// routes
	r.Get("/", handlers.UploadHandler)
	r.Post("/upload", handlers.Upload)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)

	r.Route("/download", func(r chi.Router) {
		r.Get("/{url}", handlers.DownloadHandler)
		r.Post("/{url}", handlers.Download)
	})
	http.ListenAndServe(":3001", r)
}

//Fileserver to serve static files
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

// startHTTP starts the HTTP server.
func (s *Server) startHTTP() {
	s.router = chi.NewRouter()

	// Middleware Stack
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.DefaultCompress)
	s.Router.Use(middleware.RedirectSlashes)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.Router.Use(middleware.Timeout(15 * time.Second))

	FileServer(s.router, "/assets")

	s.routes()

	s.http = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      s.router,
		Addr:         s.cfg.Port,
	}
	go func() { s.log.Fatal(s.http.ListenAndServe()) }()
}

func (s *Server) routes() {
	// Add Routes
	s.Router.Get("/", s.static())

	// Need to disable RedirectSlashes middleware to enable this
	// redirects to /debug/prof/ which causes redirect loop
	//s.Router.Mount("/debug", middleware.Profiler())

	// Add API routes
	s.router.Mount("/api", s.apiRoutes())

	// Handle NotFound
	s.router.NotFound(s.notFound())
}

func (s *Server) apiRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/thing/{id}", s.getThing())

	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.FileServer(data.Assets)

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		fs.ServeHTTP(w, r)
	}))
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panicf("Failed to encode JSON: %v\n", err)
	}
}

func writeResp(w http.ResponseWriter, data interface{}) {
	writeJSON(w, data, http.StatusOK)
}

func writeError(w http.ResponseWriter, err error, status int) {
	writeJSON(w, map[string]string{
		"Error": err.Error(),
	}, status)
}

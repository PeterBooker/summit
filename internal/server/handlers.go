package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type errResponse struct {
	Code string `json:"code,omitempty"`
	Err  string `json:"error"`
}

func (s *Server) static() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		w.Write([]byte(`Index Page`))
	}
}

func (s *Server) notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte(`Not Found`))
	}
}

// getThing fetches the data for a thing by ID
func (s *Server) getThing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if thingID := chi.URLParam(r, "id"); thingID != "" {
			// Get Thing with ID
			thing := ""
			writeResp(w, thing)
		} else {
			var resp errResponse
			resp.Err = "You must specify a valid Thing ID."
			writeResp(w, resp)
		}
	}
}

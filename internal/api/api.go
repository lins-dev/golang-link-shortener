package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

type PostBody struct {
	URL string `json:"url,omitempty"`
}

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.RequestID)
		r.Use(middleware.Recoverer)
		r.Use(middleware.RealIP)

		r.Route("/url", func(r chi.Router) {
			r.Get("/{code}", HandleGetShortenUrl(db))
			r.Post("/shorten", HandlePostShortUrl(db))
		})
	})

	return r
}
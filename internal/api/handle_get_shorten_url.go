package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lins-dev/golang-link-shortener.git/internal/repository"
)

func HandleGetShortenUrl(repository repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var code = chi.URLParam(r, "code")
		data, err := repository.FindFullUrl(r.Context(), code)
		if err != nil {
			SendJson(w, Response{Error: "invalid url"}, http.StatusBadRequest)
			return
		}
		SendJson(w, Response{Data: data}, http.StatusOK)
	}
}
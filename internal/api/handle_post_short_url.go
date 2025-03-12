package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/lins-dev/golang-link-shortener.git/internal/repository"
	"github.com/redis/go-redis/v9"
)

func HandlePostShortUrl(repository repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJson(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}
		if _, err := url.Parse(body.URL); err != nil {
			SendJson(w, Response{Error: "invalid URL"}, http.StatusBadRequest)
			return
		}
		
		code, err := repository.StoreShortenedUrl(r.Context(), body.URL)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				SendJson(w, Response{Error: "code not found"}, http.StatusNotFound)
				return
			}
			slog.Error("failed to store url", "error", err)
			SendJson(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}
		SendJson(w, Response{Data: code}, http.StatusCreated)
	}
}
package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleGetShortenUrl(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var code = chi.URLParam(r, "code")
		data, ok := db[code]
		if !ok {
			SendJson(w, Response{Error: "db invalid url"}, http.StatusBadRequest)
			return
		}
		SendJson(w, Response{Data: data}, http.StatusOK)
	}
}
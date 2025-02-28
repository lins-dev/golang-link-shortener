package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func HandlePostShortUrl(db map[string]string) http.HandlerFunc {
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
		
		code := GenerateCode()
		db[code] = body.URL
		SendJson(w, Response{Data: code}, http.StatusCreated)
	}
}
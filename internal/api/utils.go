package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func SendJson(w http.ResponseWriter, response Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(response)
	if err != nil {
		slog.Error("error in JSON marshal", "error", err)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error in write response", "error", err)
		return
	}
}
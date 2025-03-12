package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/lins-dev/golang-link-shortener.git/internal/api"
	"github.com/lins-dev/golang-link-shortener.git/internal/repository"
	"github.com/redis/go-redis/v9"
)

func main()  {
	if err := run(); err != nil {
		slog.Error("error in server", "error", err)
		return
	}
	slog.Info("server started")
}

func run() error {
	// db := make(map[string]string)
	connection := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	repository := repository.NewRepository(connection)
	handler := api.NewHandler(repository)
	server := http.Server{
		ReadTimeout: 10*time.Second,
		IdleTimeout: time.Second,
		WriteTimeout: 10*time.Second,
		Addr: ":8080",
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
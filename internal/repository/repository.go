package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type repository struct {
	connection *redis.Client
}

type Repository interface {
	StoreShortenedUrl(ctx context.Context, _url string) (string, error)
	FindFullUrl(ctx context.Context, code string) (string, error)
}

func NewRepository(connection *redis.Client) Repository {
	return repository{connection: connection}
}

func (r repository) StoreShortenedUrl(ctx context.Context, _url string) (string, error) {
	var code string
	for range 5 {
		code = GenerateCode()
		slog.Info("code", "code", code)
		err := r.connection.HGet(ctx, "shortener", code).Err()
		slog.Error("storeshortened", "error", err)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}
			return "", fmt.Errorf("failed to get code from shortener hashmap: %w", err)
		}

	}

	err := r.connection.HSet(ctx, "shortener", code, _url).Err()
	if err != nil {
		return "", fmt.Errorf("failed to store code in shortener hashmap: %w", err)
	}

	return code, nil
}

func (r repository) FindFullUrl(ctx context.Context, code string) (string, error) {
	fullUrl, err := r.connection.HGet(ctx, "shortener", code).Result()
	if err != nil {
		return "", fmt.Errorf("url not found: %w", err)
	}
	return fullUrl, nil
}
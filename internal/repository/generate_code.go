package repository

import "math/rand/v2"

func GenerateCode() string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const limit = 8
	bytes := make([]byte, limit)

	for i := range limit {
		bytes[i] = characters[rand.IntN(len(characters))]
	}
	return string(bytes)
}
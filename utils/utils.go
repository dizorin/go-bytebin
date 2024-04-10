package utils

import (
	"math/rand"
	"time"
)

var (
	r        = rand.New(rand.NewSource(time.Now().UnixMilli()))
	alphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// generateID generates a random alphanumeric ID of length 7
func GenerateID() string {
	b := make([]byte, 7)
	for i := range b {
		b[i] = alphaNum[r.Intn(len(alphaNum))]
	}
	return string(b)
}

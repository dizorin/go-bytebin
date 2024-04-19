package utils

import (
	"math/rand"
	"regexp"
	"time"
)

var (
	r          *rand.Rand
	alphaNum   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	RegexToken = regexp.MustCompile(`^/[a-zA-Z\d]{7}$`)
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixMilli())) //nolint:gosec
}

// GenerateID generates a random alphanumeric ID of length 7
func GenerateID() string {
	b := make([]byte, 7)
	for i := range b {
		b[i] = alphaNum[r.Intn(len(alphaNum))]
	}
	return string(b)
}

package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSlug(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

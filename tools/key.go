package tools

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken() string {
	r := make([]byte, 32)
	rand.Read(r)

	return base64.URLEncoding.EncodeToString(r)
}

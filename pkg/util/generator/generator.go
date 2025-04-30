package generator

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	return base64.RawStdEncoding.EncodeToString(salt), nil
}

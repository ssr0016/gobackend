package encrypt

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	memory     = 64 * 1024
	time       = 3
	threads    = 4
	keyLength  = 32
	saltLength = 16
)

func HashPassword(password, salt string) (string, error) {
	// Convert salt to byte slice
	saltBytes, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return "", fmt.Errorf("invalid salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), saltBytes, time, memory, threads, keyLength)

	return base64.RawStdEncoding.EncodeToString(hash), nil
}

func VerifyPassword(password, salt, hashedPassword string) (bool, error) {
	newHash, err := HashPassword(password, salt)
	if err != nil {
		return false, err
	}

	return newHash == hashedPassword, nil
}

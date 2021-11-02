package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// Trims 'password' value, hash it with SHA256 algorithm and return it as a string.
// Can throw an error when trimmed 'password' is empty.
func HashPassword(password string) (string, error) {
	password = strings.TrimSpace(password)
	if len(password) < 1 || password == "" {
		return "", errors.New("password can't be empty")
	}
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Hash 'password' value with SHA256 algorithm and return a result of comparsion with passed
// 'hash' value. If password is empty it will return 'false'.
func CheckPasswordHash(password, hash string) bool {
	passwordHash, error := HashPassword(password)
	if error != nil {
		fmt.Println("Hashing Error: ", error)
		return false
	}
	return passwordHash == hash
}

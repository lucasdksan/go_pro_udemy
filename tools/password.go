package tools

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("error in generated hash password")
	}

	return string(hashed), nil
}

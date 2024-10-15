package tools

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

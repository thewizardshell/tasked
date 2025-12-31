package utils

import (
	"net/mail"
	"strings"
)

func ValidateEmail(email string) bool {
	if strings.TrimSpace(email) == "" {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

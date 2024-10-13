package validate

import (
	"errors"
	"net/mail"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return errors.New("error parsing email address")
	}

	return nil
}

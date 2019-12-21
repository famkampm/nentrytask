package helper

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashingPassword(plain_password string) (string, error) {
	hashedPassword, err := Hash(plain_password)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Validate(action, username, password string) error {
	if strings.ToLower(action) == "login" || strings.ToLower(action) == "register" {
		if username == "" {
			return errors.New("Required Username")
		}
		if password == "" {
			return errors.New("Required Password")
		}
	}
	return nil
}
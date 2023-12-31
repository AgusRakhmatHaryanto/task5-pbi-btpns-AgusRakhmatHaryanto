package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string ) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(inputPass, dbPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(inputPass), []byte(dbPass))
	return err == nil
}
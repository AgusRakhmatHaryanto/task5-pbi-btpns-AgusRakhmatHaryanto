package helpers

import (
	"errors"

	"time"

	"github.com/dgrijalva/jwt-go"

)

func GenerateToken(id int, email string) (string, error) {
	var mySigningKey = []byte("AllYourBase")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	mySigningKey := []byte("AllYourBase")
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
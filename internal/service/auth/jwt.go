package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(login, secretWord string) (string, error) {

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretWord))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

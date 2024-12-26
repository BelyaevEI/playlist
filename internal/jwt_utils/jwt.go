package jwtutils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(login, secretWord string) (string, error) {

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

func ValidateToken(token, secretKey string) (bool, error) {

	tokenParse, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := tokenParse.Claims.(*Claims); ok && tokenParse.Valid {
		return true, nil
	}

	return false, fmt.Errorf("invalid token")
}

package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte("Pn1OhLGsrEluhlXMVsiu3JIItAPOh3Arzal9tFhD8aB2Mse4MNEjs0hG/F6TfHD4PQt1pg3MTLt58Z9+re9nldJf1BXsHd7XZ1auVEx0T5URAxDG7XhTE3iXGAnfp0loFndLPmSFsFDJeaLvakj8s2LIjbmUindq2fs3xkz44aBS+ckGfstiAS7gv8TCFg2J4cIC2hfacQkRZ96OXTwU6iv3mNIJ7J41HdUDow==")

type authClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func generateToken(user User) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expiresAt,
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateToken(tokenString string) (uint, string, error) {
	var claims authClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	id := claims.UserID
	username := claims.Subject
	return id, username, nil
}

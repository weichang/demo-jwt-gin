package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/weichang/demo-jwt-gin/model"
	"os"
	"time"
)

var SecretKey = []byte(os.Getenv("JWT_KEY"))

type authClaims struct {
	UserID uint `json:"userid"`
	jwt.StandardClaims
}

func GenToken(user model.User) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	c := authClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expiresAt,
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Choose specific Signature
	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (uint, string, error) {
	var claims authClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
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

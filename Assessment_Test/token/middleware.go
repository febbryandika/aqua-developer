package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GenerateJWT(username string, isAdmin bool) (string, error) {
	claims := &JWTCustomClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

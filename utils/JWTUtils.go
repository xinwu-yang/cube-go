package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func JwtSign(username string, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)),
	})
	tokenStr, _ := token.SignedString(password)
	return tokenStr
}

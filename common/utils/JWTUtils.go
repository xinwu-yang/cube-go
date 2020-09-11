package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func JwtSign(username string, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour),
	})
	tokenStr, err := token.SignedString([]byte(password))
	if err != nil {
		logs.Error(err.Error())
	}
	return tokenStr
}

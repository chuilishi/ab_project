package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	foo string `json:"foo"`
	jwt.StandardClaims
}

var claims = MyCustomClaims{
	"yjddb",
	jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 60,
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    "yjddb",
	},
}

func InitJWT() {

}

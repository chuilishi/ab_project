package service

import (
	"ab_project/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyCustomClaims jwt相关配置
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

// jwt相关配置
var claims = MyCustomClaims{
	"yjddb",
	jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 60,
		ExpiresAt: time.Now().Unix() + 60*60 - 60,
		Issuer:    "yjddb",
	},
}

// GiveJWT jwt通行证
func GiveJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(global.JWTKey))
	return ss, err

}

package service

import (
	"ab_project/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyCustomClaims jwt相关配置
type MyCustomClaims struct {
	UserName string
	jwt.StandardClaims
}

func (b *MyCustomClaims) Valid() error {
	return nil
}

// GiveJWT jwt通行证
func GiveJWT(username string) (string, error) {
	// jwt相关配置
	var claims = MyCustomClaims{
		username,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*3 - 60,
			Issuer:    "yjddb",
			Subject:   "yjddb",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString([]byte(global.JWTKey))
	return ss, err
}

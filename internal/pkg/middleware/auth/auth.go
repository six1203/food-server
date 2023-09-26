package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义过期时间
const TokenExpireDuration = time.Hour * 2

type MyClaims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

func GenToken(username string, secret string) (string, error) {
	c := MyClaims{
		username,
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(TokenExpireDuration).Unix(),
			"Issuer":    "food",
		},
	}

	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString([]byte(secret))
}

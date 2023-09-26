package auth

import (
	"food-server/internal/conf"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

func GenerateToken(username string, jwtConfig *conf.JWT) (string, error) {
	//var tokenExpireDuration = int32(time.Hour/time.Second) * jwtConfig.ExpireDuration

	c := MyClaims{
		username,
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(24 * time.Hour).Unix(),
			"Issuer":    "food",
		},
	}

	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString([]byte(jwtConfig.Secret))
}

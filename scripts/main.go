package main

import (
	"fmt"
	"food-server/internal/conf"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(username string) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username":  username,
		"ExpiresAt": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	c := &conf.JWT{
		Secret: "ToIZET+HX^a4Y$U3B%k{r{Yc=z4{%bi|",
	}
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(c.Secret))

	fmt.Println(tokenString, err)

}

func main() {
	GenerateToken("dylan")
}

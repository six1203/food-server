package main

import (
	"fmt"
	"food-server/internal/conf"
	"food-server/internal/pkg/middleware/auth"
)

func main() {
	fmt.Println(auth.GenerateToken("dylan", &conf.JWT{
		Secret:         "22222",
		ExpireDuration: 1,
	}))
}

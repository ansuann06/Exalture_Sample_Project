package main

import (
	"fmt"
	"log"

	"github.com/ansuann06/user_api/config"

	"github.com/ansuann06/user_api/routers"
)

func main() {
	config.Setup()
	r := routers.Setup()
	fmt.Println(r)
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

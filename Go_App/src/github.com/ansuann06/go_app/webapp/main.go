package main

import (
	"fmt"
	"log"

	config "github.com/ansuann06/go_app/webapp/configs"
	routes "github.com/ansuann06/go_app/webapp/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Open(); err != nil {
		fmt.Println("Not Connected")

	}
	defer config.Close()
	// Init Router
	router := gin.Default()
	// // Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}

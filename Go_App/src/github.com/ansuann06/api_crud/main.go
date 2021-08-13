package main

import (
	// "net/http"

	"github.com/ansuann06/api_crud/controllers"
	"github.com/ansuann06/api_crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	// r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}

package routes

import (
	"net/http"

	// controllers "github.com/ansuann06/go_app/webapp/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// controllers.CreateUserTable()
	router.GET("/", welcome)
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcomeee To API",
	})
	return
}

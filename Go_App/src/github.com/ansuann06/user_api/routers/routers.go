package routers

import (
	"github.com/ansuann06/user_api/config"
	"github.com/ansuann06/user_api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &controllers.APIEnv{
		DB: config.GetDB(),
	}

	r.GET("", api.GetBooks)
	r.GET("/:id", api.GetBook)
	r.POST("", api.CreateBook)
	r.PUT("/:id", api.UpdateBook)
	r.DELETE("/:id", api.DeleteBook)

	return r
}

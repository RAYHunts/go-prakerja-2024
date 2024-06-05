package routers

import (
	"uk/controllers"
	"uk/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.New()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/users", controllers.GetUsers)

	photoGroup := r.Group("/photos")
	{
		photoGroup.POST("/", middleware.Auth(), controllers.UploadPhoto)
	}



	return r
}
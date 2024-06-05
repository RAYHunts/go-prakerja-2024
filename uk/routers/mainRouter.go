package routers

import (
	"uk/controllers"
	"uk/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.New()

	// r.GET("/products", controllers.GetProducts)
	// r.GET("/products/:id", controllers.GetProduct)
	// r.POST("/products", controllers.CreateProduct)
	// r.PUT("/products/:id", controllers.UpdateProduct)
	// r.DELETE("/products/:id", controllers.DeleteProduct)

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/users", controllers.GetUsers)

	productGroup := r.Group("/products")
	{
		productGroup.Use(middleware.Auth())
		productGroup.GET("/", controllers.GetProducts)
		productGroup.GET("/:id", controllers.GetProduct)
		productGroup.POST("/", controllers.CreateProduct)
		productGroup.PUT("/:id", controllers.UpdateProduct)
		productGroup.DELETE("/:id", controllers.DeleteProduct)
	}

	return r
}
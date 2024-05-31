package controllers

import (
	"fmt"
	"net/http"

	"sesi6/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := models.GetProducts()
	log(products)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Products",
		"products": products,
	})
}

func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Product",
	})
}


func CreateProduct(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Product",
	})
}

func UpdateProduct(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Update Product with id %s", id),
	})
}

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Delete Product with id %s", id),
	})
}


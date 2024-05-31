package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"sesi6/models"

	"github.com/gin-gonic/gin"
)


func GetProducts(c *gin.Context) {
	products := models.GetProducts()
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Products",
		"products": products,
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	product := models.GetProduct(idInt)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Product",
		"product": product,
	})
}

func CreateProduct(c *gin.Context) {
	var body models.Product

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status": "error",
		})
		return
	}

	product := models.CreateProduct(body)

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Product",
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	var body models.Product

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product := models.UpdateProduct(idInt, body)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Update Product with id %s", id),
		"product": product,
	})
}

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	models.DeleteProduct(idInt)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Delete Product with id %s", id),
	})
}


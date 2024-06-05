package controllers

import (
	"net/http"
	"uk/models"

	"github.com/gin-gonic/gin"
)

func UploadPhoto(c *gin.Context) {
	var body models.Photo

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	photo := models.CreatePhoto(body)

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload Photo",
		"photo":    photo,
	})
}

func GetPhotos(c *gin.Context) {
	photos := models.GetPhotos()
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Photos",
		"photos":    photos,
	})
}
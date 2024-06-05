package controllers

import (
	"net/http"

	"uk/helper"
	"uk/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	user := models.Register(body)

	c.JSON(http.StatusOK, gin.H{
		"message": "Register",
		"user":    user,
	})
}

func Login(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	user, err := models.Login(body)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"Token":   helper.GenerateToken(user.Email),
	})
}

func GetUsers(c *gin.Context) {
	users := models.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Users",
		"users":   users,
	})
}
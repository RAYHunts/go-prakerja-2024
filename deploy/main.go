package main

import (
	"sesi7/database"
	"sesi7/models"
	"sesi7/routers"
)


func main() {
	database.StartDB()
	database.GetDB().AutoMigrate(&models.Product{})
	database.GetDB().AutoMigrate(&models.User{})
	routers.StartServer().Run(":8080")
}

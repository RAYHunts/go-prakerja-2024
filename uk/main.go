package main

import (
	"uk/database"
	"uk/models"
	"uk/routers"
)


func main() {
	database.StartDB()
	database.GetDB().AutoMigrate(&models.Photo{})
	database.GetDB().AutoMigrate(&models.User{})
	routers.StartServer().Run(":8080")
}

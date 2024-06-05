package main

import (
	"uk/database"
	"uk/models"
	"uk/routers"
)


func main() {
	database.StartDB()
	database.GetDB().AutoMigrate(&models.Product{})
	database.GetDB().AutoMigrate(&models.User{})
	routers.StartServer().Run(":8080")
}

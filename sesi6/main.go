package main

import (
	"sesi6/database"
	"sesi6/routers"
)


func main() {
	database.StartDB()
	routers.StartServer().Run(":8080")
}

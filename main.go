package main

import (
	"challange-2/database"
	"challange-2/routers"
)

func main() {
	PORT := "8080"

	database.StartDB()

	routers.StartServer(PORT).Run(":" + PORT)
}

package main

import (
	// book adalah directory root project go yang kita buat
	"api-go-mysql/models"

	"api-go-mysql/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run()
}

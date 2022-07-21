package main

import (
	"api-go-mysql/models"

	"api-go-mysql/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run()
}

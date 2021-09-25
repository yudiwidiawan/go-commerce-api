package main

import (
	"fp-jcc-go-2021-commerce/models"
	"fp-jcc-go-2021-commerce/routes"
)

func main() {

	db := models.SetupDB()
	// db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run(":8080")
}

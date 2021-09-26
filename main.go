package main

import (
	"fp-jcc-go-2021-commerce/models"
	"fp-jcc-go-2021-commerce/routes"
)

// @title Commerce API by Swagger
// @version 1.0
// @description This is a sample api of eccomerce for Final Project JCC Golang 2021.
// @description All of this is original result of learning from Bootcamp, Googling, and copy-paste-ing.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerToken
func main() {

	db := models.SetupDB()

	r := routes.SetupRoutes(db)
	r.Run()
}

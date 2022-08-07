package main

import (
	"element.com/m/controllers"
	"element.com/m/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	// load environment variables from .env
	godotenv.Load()

	// set up gin
	r := gin.Default()

	// connect to database
	models.ConnectDatabase()

	// set up routes
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:address", controllers.GetUser)

	r.Run()
}

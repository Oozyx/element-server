package main

import (
	"element.com/m/controllers"
	"element.com/m/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:address", controllers.GetUser)
	r.POST("/users/increment/:address", controllers.IncrementOwned)

	r.Run()
}

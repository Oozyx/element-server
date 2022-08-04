package controllers

import (
	"net/http"

	"element.com/m/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func IncrementOwned(c *gin.Context) {
	// validate input
	input := c.Param("address")
	if (!common.IsHexAddress(input)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}	

	// find or create user
	var user models.User
	models.DB.Where(models.User{Address: common.HexToAddress(input)}).FirstOrCreate(&user)

	// increment
	user.NumOwned++
	models.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUser(c *gin.Context) {
	// validate input
	input := c.Param("address")
	if (!common.IsHexAddress(input)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}	

	// find or create user
	var user models.User
	models.DB.Where(models.User{Address: common.HexToAddress(input)}).FirstOrCreate(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
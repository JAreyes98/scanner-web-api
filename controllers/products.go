package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scanner-web-api/dbsetup"
	"scanner-web-api/models"
)

func All(c *gin.Context) {
	var data []models.Product
	err := dbsetup.DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

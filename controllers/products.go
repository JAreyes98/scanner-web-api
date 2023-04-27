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

func ProductSearch(c *gin.Context) {
	var product models.Product
	barcode := c.Param("barcode")

	if err := dbsetup.DB.Where("barcode = ?", barcode).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha encontrado el producto con codigo de barras: " + barcode})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

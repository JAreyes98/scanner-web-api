package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllProducts(c *gin.Context) {
	var data []Product
	err := DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func ProductSearch(c *gin.Context) {
	var product Product
	barcode := c.Param("barcode")

	if err := DB.Where(" \"CodBarraSupermercado\" = ?", barcode).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha encontrado el producto con codigo de barras: " + barcode})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

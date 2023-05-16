package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllPlantas(c *gin.Context) {

	var data []Planta
	err := DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
func AllCamiones(c *gin.Context) {

	var data []Camion
	err := DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
func AllBodegas(c *gin.Context) {

	var data []Bodega
	err := DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func AllOrders(c *gin.Context) {
	var data []Order
	err := DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func FindInventory(c *gin.Context) {
	warehouse := c.Query("warehouse")
	plant := c.Query("plant")

	var query string = ""

	if warehouse != "" {
		query += " bodega='" + warehouse + "'"
	}

	if plant != "" && warehouse != "" {
		query += " and Planta='" + plant + "'"
	} else if plant != "" {
		query += " Planta='" + plant + "'"
	}

	var data []Inventory
	err := DB.Where(query).Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

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

	//var data []Bodega
	//err := DB.Find(&data).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	c.JSON(http.StatusOK, gin.H{"data": Order{ID: 1}})
}

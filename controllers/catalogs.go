package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scanner-web-api/dbsetup"
	"scanner-web-api/models"
)

func AllPlantas(c *gin.Context) {

	var data []models.Planta
	err := dbsetup.DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
func AllCamiones(c *gin.Context) {

	var data []models.Camion
	err := dbsetup.DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
func AllBodegas(c *gin.Context) {

	var data []models.Bodega
	err := dbsetup.DB.Find(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scanner-web-api/dbsetup"
	"scanner-web-api/models"
)

func CheckProductionStock(context *gin.Context) {
	var prod models.ProductionInventoryDto

	if err := context.ShouldBindJSON(&prod); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var exists string
	sql := "call sp_inventario_produccion(?, ?, ?, ?); "
	if err := dbsetup.DB.Raw(sql, prod.Planta.Planta, prod.Date, prod.Barcode, prod.IdInventario).Scan(&exists); err.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": exists})
}

func CheckOrder(context *gin.Context) {
	var prod models.ChargeInventory

	if err := context.ShouldBindJSON(&prod); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var exists string
	sql := "call sp_inventario_pedidos(?, ?, ?, ?); "
	if err := dbsetup.DB.Raw(sql, prod.Date, prod.Barcode, prod.OrderId, prod.Camion.ID).Scan(&exists); err.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": exists})
}

func CheckWarehouse(context *gin.Context) {
	var prod models.WarehouseInventory

	if err := context.ShouldBindJSON(&prod); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var exists string
	sql := "call sp_inventario_Bodega(?, ?, ?, ?); "
	if err := dbsetup.DB.Raw(sql, prod.Bodega.Codigo, prod.Date, prod.Barcode, prod.IdInventario).Scan(&exists); err.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": exists})
}

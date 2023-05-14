package main

import (
	"github.com/gin-gonic/gin"
	"scanner-web-api/controllers"
	"scanner-web-api/dbsetup"
	"scanner-web-api/middleware"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.XssMiddleware())
	//Uncomment when mode is ReleaseMode
	//r.Use(middleware.SecurityMiddleware())

	dbsetup.ConnectDatabase()
	apiAuth := r.Group("/api/auth")
	apiAuth.POST("/jwt", controllers.GenerateToken)

	api := r.Group("/api/v1").Use(middleware.Auth())
	api.GET("/products", controllers.AllProducts)
	api.GET("/products/:barcode", controllers.ProductSearch)

	api.GET("/catalogs/plants", controllers.AllPlantas)
	api.GET("/catalogs/trucks", controllers.AllCamiones)
	api.GET("/catalogs/warehouse", controllers.AllBodegas)

	api.POST("/inventory/production", controllers.CheckProductionStock)
	api.POST("/inventory/order", controllers.CheckOrder)
	api.POST("/inventory/warehouse", controllers.CheckWarehouse)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

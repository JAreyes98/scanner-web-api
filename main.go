package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(CorsMiddleware())
	r.Use(XssMiddleware())
	//Uncomment when mode is ReleaseMode
	//r.Use(middleware.SecurityMiddleware())

	ConnectDatabase()
	apiAuth := r.Group("/api/auth")
	apiAuth.POST("/jwt", GenerateToken)

	api := r.Group("/api/v1").Use(Auth())
	api.GET("/products", AllProducts)
	api.GET("/products/:barcode", ProductSearch)

	api.GET("/catalogs/plants", AllPlantas)
	api.GET("/catalogs/trucks", AllCamiones)
	api.GET("/catalogs/warehouse", AllBodegas)

	api.POST("/inventory/production", CheckProductionStock)
	api.POST("/inventory/order", CheckOrder)
	api.POST("/inventory/warehouse", CheckWarehouse)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

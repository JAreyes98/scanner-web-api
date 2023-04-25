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
	v1 := r.Group("/api/v1")
	v1.GET("/products", controllers.All)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

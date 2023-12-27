package main

import (
	"github.com/gin-gonic/gin"
	"veltra.com/gin_playground/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

        r.GET("/", controllers.Index)
	r.GET("/reviews/product/:product_id", controllers.GetProduct)

	return r
}

package main

import (
	"rest-api/controller"
	"rest-api/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	db := model.SetupModels()

	// Provide db to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1 := r.Group("/v1")

	// Routes
	v1.GET("/rental", controller.FindRentals)
	v1.GET("/rental/:id", controller.FindRental)
	v1.POST("/rental", controller.CreateRental)
	v1.PATCH("/rental/:id", controller.UpdateRental)
	v1.DELETE("/rental/:id", controller.DeleteRental)

	// Run the server
	r.Run(":3000")
}
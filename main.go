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
	rental := v1.Group("/rental")
	rental.GET("", controller.FindRentals)
	rental.GET("/:id", controller.FindRental)
	rental.POST("", controller.CreateRental)
	rental.PATCH("/:id", controller.UpdateRental)
	rental.DELETE("/:id", controller.DeleteRental)

	// Run the server
	r.Run(":3000")
}
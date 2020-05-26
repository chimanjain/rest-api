package main

import (
	"rest-api/model"
	"rest-api/routes"

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

	routes.InitializeRentalRoutes(r)

	// Run the server
	r.Run(":3000")
}

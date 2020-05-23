package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/controller"
)


func InitializeRoutes(r *gin.Engine) {

	// Handle the index route
	v1 := r.Group("/v1")

	// Routes
	v1.GET("/rental", controller.FindRentals)
	v1.GET("/rental/:id", controller.FindRental)
	v1.POST("/rental", controller.CreateRental)
	v1.PATCH("/rental/:id", controller.UpdateRental)
	v1.DELETE("/rental/:id", controller.DeleteRental)
}

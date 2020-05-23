package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"rest-api/model"
	"time"
)

type CreateRentalInput struct {
	Name  string `json:"name" binding:"required"`
	OwnerName string `json:"ownerName" binding:"required"`
	Location string `json:"location" binding:"required"`
	Price string `json:"price" binding:"required"`
}

type UpdateRentalInput struct {
	Name  string `json:"name"`
	OwnerName string `json:"ownerName"`
	Location string `json:"location"`
	Price string `json:"price"`
	DateCreated string `json:"dateCreated"`
}

// GET /rental
// Find all rentals
func FindRentals(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var rentals []model.Rental
	db.Find(&rentals)

	c.JSON(http.StatusOK, gin.H{"data": rentals})
}

// GET /rental/:id
// Find a rental
func FindRental(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var rental model.Rental
	if err := db.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// POST /rental
// Create new rental
func CreateRental(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateRentalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rental
	rental := model.Rental{
		Name: input.Name,
		OwnerName: input.OwnerName,
		Location: input.Location,
		Price: input.Price,
		DateCreated: time.Now().UTC().String(),
	}
	db.Create(&rental)

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// PATCH /rental/:id
// Update a rental
func UpdateRental(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var rental model.Rental
	if err := db.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateRentalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&rental).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// DELETE /rental/:id
// Delete a rental
func DeleteRental(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var rental model.Rental
	if err := db.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&rental)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
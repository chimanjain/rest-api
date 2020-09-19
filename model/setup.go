package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupModels() *gorm.DB {
	dsn := "root:password@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Rental{})

	return db
}

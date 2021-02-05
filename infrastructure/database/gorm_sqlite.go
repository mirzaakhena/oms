package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("order.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databaseConnection *gorm.DB

func CreateConnection() *gorm.DB {
	dbConn, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("Database connection failed !")
	}

	return dbConn
}

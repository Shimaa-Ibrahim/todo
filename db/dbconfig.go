package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDBConnection() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to db")
	}
	DB = db
}

func Close() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}

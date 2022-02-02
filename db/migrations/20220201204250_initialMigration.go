package main

import (
	models "github/Shimaa-Ibrahim/todo/models"

	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20220201204250(txn *gorm.DB) {
	txn.AutoMigrate().CreateTable(&models.User{})
	txn.AutoMigrate().CreateTable(&models.Task{})
}

// Down is executed when this migration is rolled back
func Down_20220201204250(txn *gorm.DB) {
	txn.AutoMigrate().DropTable(&models.User{})
	txn.AutoMigrate().DropTable(&models.Task{})
}

package handler

import "gorm.io/gorm"

var DB *gorm.DB

func RegisterDB(db *gorm.DB) {
	DB = db
}

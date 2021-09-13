package db

import (
	"fmt"
	"log"
	"wm/sprint/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/sprint?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	db.AutoMigrate(&Sprint{})
}

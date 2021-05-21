package database

import (
	"modules/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	conn, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang"), &gorm.Config{})
	if err != nil {
		panic("Bağlantıda sorun var!!!")
	}
	DB = conn

	conn.AutoMigrate(&models.User{})
}

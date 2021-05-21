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
	//Kurulan bağlantıyı global olarak diğer package'lar ile yaplaşıyoruz
	DB = conn
	//Burada User Struct'ını Mysql'e migrate ederek user tablosunun otomatik oluşmasını sağlıyoruz
	conn.AutoMigrate(&models.User{})
}

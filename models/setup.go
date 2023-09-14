package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_sertifikasi_go"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&PenawaranSertifikasi{})
	database.AutoMigrate(&PendaftaranUjian{})
	database.AutoMigrate(&User{})

	DB = database
}

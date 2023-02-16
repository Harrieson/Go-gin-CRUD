package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gorestapi"))
	if err != nil {
		panic(err.Error())
	}
	database.AutoMigrate(&Product{})
	DB = database
}

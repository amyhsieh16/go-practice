package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/sqlserver"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(sqlserver.Open("server=localhost;user id=sa;password=123456;database=test"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&User{})
	DB = database
}
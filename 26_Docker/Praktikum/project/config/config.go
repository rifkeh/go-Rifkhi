package config

import (
	"fmt"
	"project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "rifkhi",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "shop",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})
}

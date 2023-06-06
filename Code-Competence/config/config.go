package config

import (
	"fmt"
	"log"
	"os"

	"competence/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	err error
)

func InitDB() error{
	errenv := godotenv.Load()

	if errenv != nil{
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed to connect to database!")
	}
	
	InitMigrate()
	return nil
}

func InitMigrate(){
	DB.AutoMigrate(&model.Items{})
	DB.AutoMigrate(&model.Admin{})
}
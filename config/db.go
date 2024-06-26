package config

import (
	"fmt"
	"os"
	"user/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func ConnectToDB() (*gorm.DB, error) {
  err := godotenv.Load(".env")
 if err != nil{
  fmt.Printf("Error loading .env file: %q", err)
 }
  dsn := os.Getenv("DSN")

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    fmt.Printf("DATABASE CONNECTION FAILE. Error: %q", err)
    return nil, err
  }

  db.AutoMigrate(&models.User{})

  DB = db 
  fmt.Println("DATABASE CONNECTED")
  return nil, err
}

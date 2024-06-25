package database
// db "github.com/fs0414/go_hobby/internal/infrastructure/database"

import (
  "os"
  "fmt"

  "github.com/joho/godotenv"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"

  "github.com/fs0414/go_hobby/internal/infrastructure/model"
)

var db *gorm.DB

func DbInit() {
  LoadEnv()

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
    os.Getenv("POSTGRES_HOST"),
    os.Getenv("POSTGRES_USER"),
    os.Getenv("POSTGRES_PASSWORD"),
    os.Getenv("POSTGRES_DB"),
    os.Getenv("POSTGRES_POST"),
    os.Getenv("POSTGRES_SSLMODE"),
    os.Getenv("POSTGRES_TIMEZONE"),
  )

  var err error
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
      panic("failed to connect to database")
  }

  db.AutoMigrate(&model.User{}, &model.Board{})
}

func LoadEnv() {
	err := godotenv.Load(".env")
	
	if err != nil {
		fmt.Printf("not loaded env: %v", err)
	}
}

func GetDb() *gorm.DB {
  return db
}
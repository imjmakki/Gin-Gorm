package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

func DbConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load the environment file")
	}
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASS")
	Host := os.Getenv("DB_HOST")
	Name := os.Getenv("DB_NAME")
}

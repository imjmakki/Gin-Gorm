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
}

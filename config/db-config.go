package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func DbConnect() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load the environment file")
	}
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASS")
	Host := os.Getenv("DB_HOST")
	Name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:8889)/%s?charset=utf8&parseTime=True&loc=local", User, Password, Host, Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}
	//db.AutoMigrate()
	return db
}

func CloseConnect(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to Close the Connection")
	}
	dbSQL.Close()
}

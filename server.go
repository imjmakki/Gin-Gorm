package main

import (
	"GoRm-Gin/API"
	"GoRm-Gin/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB    = config.DbConnect()
	authApi API.AuthApi = API.NewAuthApi()
)

func main() {
	defer config.CloseConnect(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login")
		authRoutes.POST("/register")
	}

	r.Run()
}

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
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login")
	}

	r.Run()
}

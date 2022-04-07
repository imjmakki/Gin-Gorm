package API

import "github.com/gin-gonic/gin"

type AuthApi interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type AuthApi struct {
}

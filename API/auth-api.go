package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthApi interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authApi struct {
}

func NewAuthApi() AuthApi {
	return &authApi{}
}

func (c *authApi) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You are logged in",
	})
}

func (c *authApi) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You've been added to the system now",
	})
}

package API

import "github.com/gin-gonic/gin"

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

}

func (c *authApi) Register(ctx *gin.Context) {

}

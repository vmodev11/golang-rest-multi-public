package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/cmd/entity-server/api"
)

func initAuthRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("auth")
	{
		AuthRouter.POST("login", api.Login)       //login
		AuthRouter.POST("register", api.Register) //register
	}

}

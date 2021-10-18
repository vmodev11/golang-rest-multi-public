package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter(env string) *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.Cors())
	urlPrefix := fmt.Sprintf("%s/api/v1", env)
	ApiGroup := Router.Group(urlPrefix)

	initAuthRouter(ApiGroup)
	return Router
}

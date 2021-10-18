package utils

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/cmd/entity-server/config"
)

func Cors() gin.HandlerFunc {

	return cors.New(cors.Config{
		AllowOrigins:     config.GetConfig().AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		/* 	AllowOriginFunc: func(origin string) bool {
			return origin == "*" // using pattern
		}, */
		MaxAge: 12 * time.Hour,
	})
}

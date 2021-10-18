package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/cmd/entity-server/config"
	"github.com/golang/cmd/entity-server/models"
	"github.com/golang/cmd/entity-server/router"
)

func main() {
	//get env param
	env := os.Getenv("GO_ENV")
	if env == "" {
		os.Setenv("GO_ENV", "local")
		env = "local"
	}
	log.Printf("Env set to ** %s **", env)
	//get port param
	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "8080")
		port = "8080"
	}
	log.Printf("Port set to ** %s **", port)

	basePath, _ := os.Getwd()
	config.InitFromFile("config/config.toml", basePath)
	closeFunc, _ := models.InitFromSQLLite(config.GetConfig().DbConnection)
	routersInit := router.InitRouter(env)
	readTimeout := time.Minute
	writeTimeout := time.Minute
	endPoint := fmt.Sprintf(":%s", port) //config.GetConfig().ServerPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()
	defer closeFunc()
}

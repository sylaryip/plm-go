package main

import (
	"fmt"
	"go-plm/libs/config"
	"go-plm/libs/db"
	"go-plm/libs/routers"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	configPath := "config.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}
	config.InitViperConfig(configPath)
	db.InitDB()
	router = routers.InitRouter()
}

func main() {
	appConfig := config.GetServerConfig()
	err := router.Run(fmt.Sprintf(":%d", appConfig["port"]))
	if err != nil {
		panic(err)
	}
}

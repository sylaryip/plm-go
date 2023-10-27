package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-plm/libs/config"
	"go-plm/libs/db"
	"go-plm/libs/routers"
)

var router *gin.Engine

func init() {
	config.InitViperConfig()
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

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-plm/libs/config"
	"go-plm/libs/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(fmt.Sprint(config.GetServerConfig()["mode"]))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"pong": "Hello world!"})
	})

	material := r.Group("/material")
	{
		material.GET("/fuzzy/search", api.FuzzyMaterialDetail)
	}
	return r
}

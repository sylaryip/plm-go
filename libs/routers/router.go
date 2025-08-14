package routers

import (
	"fmt"
	"go-plm/libs/config"
	"go-plm/libs/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

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

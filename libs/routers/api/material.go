package api

import "github.com/gin-gonic/gin"

func FuzzyMaterialDetail(c *gin.Context) {
	c.JSON(200, gin.H{
		"Fuzzy": "OK",
	})
}

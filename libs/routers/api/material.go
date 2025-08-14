package api

import (
	"go-plm/libs/db"
	"go-plm/libs/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FuzzyMaterialDetail(c *gin.Context) {
	materialCode := c.Query("material_code")
	if materialCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "material_code is required",
		})
		return
	}

	var materialData models.MaterialData
	result := db.DB.Where("material_code = ?", materialCode).First(&materialData)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "material not found",
		})
		return
	}

	c.JSON(http.StatusOK,
		materialData,
	)
}

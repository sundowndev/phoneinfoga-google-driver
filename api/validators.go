package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type scanURL struct {
	Number uint `uri:"number" binding:"required,min=1,max=999679368229"`
}

// ValidateScanURL validates scan URLs
func ValidateScanURL(c *gin.Context) {
	var v scanURL

	if err := c.ShouldBindUri(&v); err != nil {
		errorHandling(c, err.Error())
	}
}

func errorHandling(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": msg})
	c.Abort()
}

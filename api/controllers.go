package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sundowndev/phoneinfoga/scanners"
)

func getAllNumbers(c *gin.Context) {
	c.JSON(200, gin.H{
		"numbers": []scanners.Number{},
	})
}

func localScan(c *gin.Context) {
	number := c.Param("number")
	localScan, err := scanners.LocalScan(number)

	if err != nil {
		c.JSON(500, errorResponse)
		return
	}

	c.JSON(200, gin.H{
		"result": localScan,
	})
}

func healthHandler(c *gin.Context) {
	c.JSON(200, successResponse())
}

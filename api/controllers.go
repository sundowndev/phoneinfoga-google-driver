package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sundowndev/phoneinfoga/pkg/scanners"
	"github.com/sundowndev/phoneinfoga/pkg/utils"
)

func getAllNumbers(c *gin.Context) {
	c.JSON(200, gin.H{
		"numbers": []scanners.Number{},
	})
}

func localScan(c *gin.Context) {
	number := c.Param("number")

	number = utils.FormatNumber(number)
	result, err := scanners.LocalScan(number)

	if err != nil {
		c.JSON(500, errorResponse("The number is not valid"))
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func numverifyScan(c *gin.Context) {
	number := c.Param("number")

	number = utils.FormatNumber(number)
	result, err := scanners.NumverifyScan(number)

	if err != nil {
		c.JSON(500, errorResponse())
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func googleSearchScan(c *gin.Context) {
	number := c.Param("number")

	number = utils.FormatNumber(number)
	n, err := scanners.LocalScan(number)

	if err != nil {
		c.JSON(500, errorResponse("The number is not valid"))
		return
	}

	result := scanners.GoogleSearchScan(n)

	c.JSON(200, gin.H{
		"result": result,
	})
}

func healthHandler(c *gin.Context) {
	c.JSON(200, successResponse())
}

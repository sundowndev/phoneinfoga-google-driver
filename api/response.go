package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func successResponse(msg ...string) map[string]interface{} {
	var message string

	if len(msg) == 0 {
		message = "An error occurred"
	} else {
		message = strings.Join(msg, "")
	}

	return gin.H{
		"success": true,
		"message": message,
	}
}

func errorResponse(msg ...string) map[string]interface{} {
	var message string

	if len(msg) == 0 {
		message = "An error occurred"
	} else {
		message = strings.Join(msg, "")
	}

	return gin.H{
		"success": false,
		"message": message,
	}
}

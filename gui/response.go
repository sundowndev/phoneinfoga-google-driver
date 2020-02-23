package gui

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func successResponse() map[string]interface{} {
	return gin.H{
		"success": true,
		"message": nil,
	}
}

func errorResponse(msg ...string) map[string]interface{} {
	var message string

	if len(msg) == 0 {
		message = "An error occured"
	} else {
		message = strings.Join(msg, "")
	}

	return gin.H{
		"success": false,
		"message": message,
	}
}

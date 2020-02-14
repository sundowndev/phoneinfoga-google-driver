package gui

import "github.com/gin-gonic/gin"

func successResponse() map[string]interface{} {
	return gin.H{
		"success": true,
		"message": nil,
	}
}

func errorResponse() map[string]interface{} {
	return gin.H{
		"success": false,
		"message": "An error occured",
	}
}

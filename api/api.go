package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Serve launches the web client
// Using Gin & Vue.js
func Serve(port int) {
	httpPort := ":" + strconv.Itoa(port)

	router := gin.Default()

	router.Group("/api").
		GET("/", healthHandler).
		GET("/numbers", getAllNumbers).
		GET("/numbers/:number/scan/local", localScan)

	router.Run(httpPort)
}

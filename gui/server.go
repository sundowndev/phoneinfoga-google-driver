package api

import (
	"net/http"
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
		GET("/numbers/:number/scan/local", localScan).
		GET("/numbers/:number/scan/numverify", numverifyScan)

	router.StaticFS("/*", http.Dir("client"))

	router.Run(httpPort)
}

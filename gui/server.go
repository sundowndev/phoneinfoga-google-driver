package api

import (
	"net/http"
	"os"
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

	dir, _ := os.Getwd()
	assetsPath := dir + "/gui/client/dist"

	router.Static("/js", assetsPath+"/js")
	router.Static("/css", assetsPath+"/css")
	router.Static("/img", assetsPath+"/img")
	router.StaticFile("/favicon.ico", assetsPath+"/favicon.ico")
	router.LoadHTMLFiles(assetsPath + "/index.html")

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run(httpPort)
}

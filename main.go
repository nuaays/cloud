package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World"})
}

func HelloHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{"message": "Hello " + name})
}

func main() {
	gin.SetMode(gin.DebugMode) //gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.GET("/:name", HelloHandler)
	// router.Run(":5000")
	http.ListenAndServe(":5000", router)
}

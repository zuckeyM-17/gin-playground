package main

import (
	"net/http"

	"app/renderer"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = renderer.CreateRenderer()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "this is title",
		})
	})

	auth := r.Group("/auth")
	auth.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.Run(":1234")
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "this is title",
		})
	})
}

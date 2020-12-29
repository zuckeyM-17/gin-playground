package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexRoutes(r *gin.Engine) {
	r.GET("/", getIndex)
	r.GET("/sign_in", getSignin)
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "this is title",
	})
}

func getSignin(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{
		"title": "this is sign_in page",
	})
}

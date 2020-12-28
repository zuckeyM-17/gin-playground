package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	auth.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

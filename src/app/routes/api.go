package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiRoutes(r *gin.Engine) {

	api := r.Group("/api")
	api.POST("/sign_in", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

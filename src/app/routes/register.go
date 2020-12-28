package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	indexRoutes(r)
	apiRoutes(r)
}

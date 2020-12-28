package main

import (
	"app/renderer"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = renderer.CreateRenderer()
	routes.RegisterRoutes(r)
	r.Run(":1234")
}

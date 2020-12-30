package main

import (
	"app/middleware"
	"app/renderer"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.SessionStore())
	r.Use(middleware.CsrfProtection())

	r.Static("/assets", "./assets")

	r.HTMLRender = renderer.CreateRenderer()
	routes.RegisterRoutes(r)
	r.Run(":1234")
}

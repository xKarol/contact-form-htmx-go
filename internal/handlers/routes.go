package handlers

import (
	"app/internal/templates/pages"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		templ.Handler(pages.HomePage("Home")).Component.Render(c, c.Writer)
	})

	apiGroup := router.Group("/api")
	apiGroup.POST("/health-check", HealthCheck)
	apiGroup.POST("/contact", CreateContact)

	router.NoRoute(func(c *gin.Context) {
		templ.Handler(pages.ErrorPage("404", "Not found")).Component.Render(c, c.Writer)
	})
}

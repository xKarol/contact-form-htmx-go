package handlers

import (
	"app/internal/templates/pages"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		templ.Handler(pages.HomePage("Home")).Component.Render(c, c.Writer)
	})

	router.POST("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	router.POST("/contact", CreateContact)

	router.NoRoute(func(c *gin.Context) {
		templ.Handler(pages.ErrorPage("404", "Not found")).Component.Render(c, c.Writer)
	})
}

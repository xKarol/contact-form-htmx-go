package main

import (
	"app/internal/routes"
	"app/internal/templates/pages"
	"app/internal/utils"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	router := gin.Default()

	router.Static("/static", "./dist")

	router.GET("/", func(c *gin.Context) {
		templ.Handler(pages.HomePage("Home")).Component.Render(c, c.Writer)
	})

	routes.Init(router)

	router.Run()
}

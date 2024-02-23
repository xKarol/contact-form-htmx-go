package main

import (
	"app/routes"
	"app/templates/pages"
	"app/utils"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		templ.Handler(pages.HomePage("Home")).Component.Render(c, c.Writer)
	})

	routes.Init(router)

	router.Run()
}

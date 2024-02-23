package main

import (
	"app/internal/handlers"
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

	handlers.Init(router)

	router.Run()
}

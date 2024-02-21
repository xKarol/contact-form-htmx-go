package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/contact", func(c *gin.Context) {
		fields := []string{"email", "firstName", "lastName", "message"}

		for _, field := range fields {
			fmt.Println(c.Request.FormValue(field))
		}
		c.Status(http.StatusOK)
	})

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	router.Run()
}

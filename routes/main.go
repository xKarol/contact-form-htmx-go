package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.POST("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	router.POST("/contact", func(c *gin.Context) {
		fields := []string{"email", "firstName", "lastName", "message"}

		for _, field := range fields {
			fmt.Println(c.Request.FormValue(field))
		}
		c.Status(http.StatusOK)
	})
}

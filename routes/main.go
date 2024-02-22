package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Init(router *gin.Engine) {
	router.POST("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	router.POST("/contact", func(c *gin.Context) {

		firstName := c.Request.FormValue("firstName")
		lastName := c.Request.FormValue("lastName")
		email := c.Request.FormValue("email")
		message := c.Request.FormValue("message")

		from := mail.NewEmail("Example User", "test@example.com")
		subject := "Sending with Twilio SendGrid is Fun"
		to := mail.NewEmail("Example User", email)
		content := fmt.Sprintf("First Name: %s, Last Name: %s, Email: %s, Message: %s", firstName, lastName, email, message)

		htmlContent := "<strong>" + content + "</strong>"
		sendMail := mail.NewSingleEmail(from, subject, to, content, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

		if gin.Mode() != "release" {
			client.BaseURL = "http://127.0.0.1:3030/v3/mail/send"
		}

		_, err := client.Send(sendMail)
		if err != nil {
			log.Println("err:", err)
		}

		c.Status(http.StatusOK)
	})
}

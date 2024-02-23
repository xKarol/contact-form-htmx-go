package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/net/html"
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

		fmt.Println(message)

		dir, _ := os.Getwd()
		templatePath := path.Join(dir, "/internal/templates/contact-confirm.html")
		emailTemplate, err1 := template.ParseFiles(templatePath)
		if err1 != nil {
			log.Println("Error parsing email template:", err1)
			c.Status(http.StatusInternalServerError)
			return
		}

		var emailBodyBuffer bytes.Buffer
		if err2 := emailTemplate.Execute(&emailBodyBuffer, map[string]string{
			"firstName": firstName,
			"lastName":  lastName,
		}); err2 != nil {
			log.Println("Error executing email template:", err2)
			c.Status(http.StatusInternalServerError)
			return
		}

		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		if gin.Mode() != "release" {
			client.BaseURL = "http://127.0.0.1:3030/v3/mail/send"
		}

		from := mail.NewEmail("Example User", "test@example.com")
		subject := "[Company Name] Your Contact Submission was Received"
		htmlContent := emailBodyBuffer.String()
		to := mail.NewEmail(firstName+" "+lastName, email)
		content, _ := extractTextFromHTML(htmlContent)
		sendMail := mail.NewSingleEmail(from, subject, to, content, htmlContent)
		_, err := client.Send(sendMail)
		if err != nil {
			log.Println("err:", err)
		}

		c.Status(http.StatusOK)
	})
}

func extractText(node *html.Node) string {
	var result string

	if node.Type == html.TextNode {
		result += strings.TrimSpace(node.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result += extractText(child)
	}

	return result
}

func extractTextFromHTML(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	return extractText(doc), nil
}

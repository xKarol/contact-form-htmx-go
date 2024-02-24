package handlers

import (
	"app/internal/templates/components"
	"app/internal/templates/mails"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/net/html"
)

type Contact struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required,max=30"`
	LastName  string `form:"lastName" json:"lastName" binding:"required,max=30"`
	Email     string `form:"email" json:"email" binding:"required,email"`
	Message   string `form:"message" json:"message" binding:"required,min=5,max=500"`
}

func CreateContact(c *gin.Context) {
	var form Contact
	if err := c.ShouldBind(&form); err != nil {
		templ.Handler(components.Alert("error", "Validation Error:"+err.Error())).Component.Render(c, c.Writer)
		return
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	if gin.Mode() != "release" {
		client.BaseURL = "http://127.0.0.1:3030/v3/mail/send"
	}

	from := mail.NewEmail("Example User", "test@example.com")
	subject := "[Company Name] Your Contact Submission was Received"
	to := mail.NewEmail(form.FirstName+" "+form.LastName, form.Email)

	templContent, err := templ.ToGoHTML(c, mails.ContactConfirm(form.FirstName, form.LastName))
	if err != nil {
		log.Println("err:", err)

		templ.Handler(components.Alert("error", "Something went wrong...")).Component.Render(c, c.Writer)
		return
	}

	htmlContent := fmt.Sprintf("%s", templContent)
	content, _ := extractTextFromHTML(htmlContent)

	sendMail := mail.NewSingleEmail(from, subject, to, content, htmlContent)
	_, err1 := client.Send(sendMail)
	if err1 != nil {
		log.Println("err:", err1)
		templ.Handler(components.Alert("error", "Something went wrong...")).Component.Render(c, c.Writer)
		return
	}
	templ.Handler(components.Alert("info", "Email has been sent.")).Component.Render(c, c.Writer)
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

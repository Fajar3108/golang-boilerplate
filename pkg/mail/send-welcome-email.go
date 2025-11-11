package mail

import (
	_ "embed"
	"text/template"

	"github.com/gofiber/fiber/v2/log"
)

//go:embed template/welcome.html
var welcomeEmailTemplate string

func SendWelcomeEmail(userName string, userEmail string) {
	tmpl, err := template.New("welcome").Parse(welcomeEmailTemplate)

	if err != nil {
		log.Errorf("Failed to parse embedded template: %v", err)
		return
	}

	data := map[string]string{
		"Name": userName,
	}

	if err := SendMailWithTemplate(
		userEmail,
		"Welcome to Online Course Platform",
		tmpl,
		data,
	); err != nil {
		log.Errorf("Failed to send welcome email to %s: %v", userEmail, err)
	}
}

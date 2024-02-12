package mail

import (
	"github.com/ArthurMaverick/zap-ai/pkg/env"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendGridMail(name, email, subject, fileName, token string) (*rest.Response, error) {
	from := mail.NewEmail("admin", "arthuracs18@gmail.com")
	to := mail.NewEmail(name, email)
	subjectMail := subject
	template := ParseHtml(fileName, map[string]string{
		"to":    email,
		"token": token,
	})
	sendgridAPIKey, err := env.GodoEnv("SENDGRID_API_KEY")
	if err != nil {
		return nil, err
	}

	message := mail.NewSingleEmail(from, subjectMail, to, template, template)
	client := sendgrid.NewSendClient(sendgridAPIKey)
	return client.Send(message)
}

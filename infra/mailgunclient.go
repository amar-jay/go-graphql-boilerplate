package mailgunclient

import (
	"context"
	"errors"

	"github.com/amar-jay/go-api-boilerplate/config"
	"github.com/mailgun/mailgun-go/v4"
)

type MailgunClient interface {
	Welcome(subject, text, to, htmlStr string) error
	ResetPassword(subject, text, to, htmlStr string) error
}

type mailgunClient struct {
	conf   config.Config
	client *mailgun.MailgunImpl
}

func NewMailgunClient(c config.Config) *mailgunClient {
	mg_conf := config.GetMailgunConfig()
	return &mailgunClient{conf: c, client: mailgun.NewMailgun(
		mg_conf.Domain,
		mg_conf.APIKey,
	)}
}
func SendWelcomeEmail(toEmail string) error {
	mg_conf := config.GetMailgunConfig()
	mg := mailgun.NewMailgun(mg_conf.Domain, mg_conf.APIKey)
	html_template := "<html><body><h1>Welcome to the API</h1><p>Thanks for signing up</p></body></html>"
	message := mg.NewMessage(
		mg_conf.From,
		"Welcome to the API",
		html_template,
		toEmail,
	)
	context := context.Background()
	_, _, err := mg.Send(context, message)
	if err != nil {
		return errors.New("Error sending welcome email")
	}
	return nil
}
func SendResetPasswordEmail(toEmail string, token string) error {
	mg_conf := config.GetMailgunConfig()
	mg := mailgun.NewMailgun(mg_conf.Domain, mg_conf.APIKey)
	html_template := "<html><body><h1>Reset Password</h1><p>Click on the link below to reset your password</p><a href='http://localhost:8080/reset-password?token=" + token + "'>Reset Password</a></body></html>"
	message := mg.NewMessage(
		mg_conf.From,
		"Reset Password",
		html_template,
		toEmail,
	)
	context := context.Background()
	_, _, err := mg.Send(context, message)
	if err != nil {
		return errors.New("Error sending reset password email")
	}
	return nil
}

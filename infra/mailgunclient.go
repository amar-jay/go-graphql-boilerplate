package mailgunclient

import (
	"github.com/amar-jay/go-api-boilerplate/config"
	"github.com/mailgun/mailgun-go/v4"
)



type MailgunClient interface {
  Welcome(subject, text, to, htmlStr string) error
  ResetPassword(subject, text, to, htmlStr string) error
}


type mailgunClient struct {
  conf config.Config
  client *mailgun.MailgunImpl
}

func NewMailgunClient(c config.Config) *mailgunClient {
  mg_conf := config.GetMailgunConfig()
  return &mailgunClient{conf: c, client: mailgun.NewMailgun(
    mg_conf.Domain,
    mg_conf.APIKey,
    )}
}

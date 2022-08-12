package mailgunclient

import (
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
  return &mailgunclient{}
}

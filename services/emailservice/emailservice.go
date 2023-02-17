package emailservice

import (
	"errors"

	"github.com/amar-jay/go_api_boilerplate/infra/mailgunclient"
)

// "github.com/amar-jay/go_api_boilerplate/infra/mailgunclient"
type EmailService interface {
	Welcome(toEmail string) error
	ResetPassword(toEmail string, token string) error
}

type emailService struct {
}

// NewEmailService returns a new instance of the email service
func NewEmailService() EmailService {
	return &emailService{}
}

// Welcome email
func (s *emailService) Welcome(toEmail string) error {
	if err := mailgunclient.SendWelcomeEmail(toEmail); err != nil {
		return errors.New("Error sending welcome email")
	}
	return nil
}

// resetPassword
func (s *emailService) ResetPassword(toEmail, token string) error {
	if err := mailgunclient.SendResetPasswordEmail(toEmail, token); err != nil {
		return errors.New("Error sending reset password email")
	}
	return nil
}

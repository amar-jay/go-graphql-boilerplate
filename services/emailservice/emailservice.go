package emailservice

import "errors"

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

//Welcome email
func (s *emailService) Welcome(toEmail string) error {
	return errors.New("Not implemented")
}


// resetPassword
func (s *emailService) ResetPassword(toEmail, token string) error {
	return nil
}

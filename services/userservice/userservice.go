package userservice

import "github.com/amar-jay/go-api-boilerplate/domain/user"

type UserService interface {
		Login()
		Register()
		ResetPassword()
		ForgotPassword()
		GetUserByID(id uint) (*user.User, error)
}

type userService struct {
	pepper string
}




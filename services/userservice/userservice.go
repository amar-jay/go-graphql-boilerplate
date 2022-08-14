package userservice

import (
	"github.com/amar-jay/go-api-boilerplate/domain/user"
	rdms "github.com/amar-jay/go-api-boilerplate/common/randomstring"
	"github.com/amar-jay/go-api-boilerplate/common/hmachash"
	pswd_repo "github.com/amar-jay/go-api-boilerplate/repositories/password_reset"

	"github.com/amar-jay/go-api-boilerplate/repositories/user_repo"
)

type UserService interface {
		Register(user *user.User) error
		Update(user *user.User) error
		GetUserByID(id uint) (*user.User, error)
}

type userService struct {
	pepper string
	repo user_repo.Repo
	pswd pswd_repo.Repo
	rs rdms.RandomString
	hmac hmachash.HMAC 
}

func NewUserService(repo user_repo.Repo, pswd pswd_repo.Repo, rs rdms.RandomString, hmac hmachash.HMAC, pepper string ) UserService {

	return &userService{
		repo: repo,
		pepper: pepper,
		pswd: pswd,
		rs: rs,
		hmac: hmac,
	}
}

func (us *userService) Register(u *user.User) error {
	return nil
}

func (us *userService) Update(u *user.User) error {
	return nil
}


func (us *userService) GetUserByID(id uint) (*user.User, error) {

	return nil, nil
}


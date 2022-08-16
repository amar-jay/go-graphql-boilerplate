package userservice

import (
	"regexp"

	"github.com/amar-jay/go-api-boilerplate/common/hmachash"
	rdms "github.com/amar-jay/go-api-boilerplate/common/randomstring"
	"github.com/amar-jay/go-api-boilerplate/domain/user"
	pswd_repo "github.com/amar-jay/go-api-boilerplate/repositories/password_reset"
	"github.com/amar-jay/go-api-boilerplate/repositories/user_repo"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
		ComparePassword(inputpswd string, dbpswd string) error
		Register(user *user.User) error
		Update(user *user.User) error
		GetUserByID(id uint) (*user.User, error)
		GetUserByEmail(email string) (*user.User, error)
		GetUsers() ([]*user.User, error)
}

type userService struct {
	pepper string
	Repo user_repo.Repo
	pswd pswd_repo.Repo
	rs rdms.RandomString
	hmac hmachash.HMAC 
}

func NewUserService(repo user_repo.Repo, pswd pswd_repo.Repo, rs rdms.RandomString, hmac hmachash.HMAC, pepper string ) UserService {

	return &userService{
		Repo: repo,
		pepper: pepper,
		pswd: pswd,
		rs: rs,
		hmac: hmac,
	}
}

func (us *userService) Register(u *user.User) error {
	hashed, err := us.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashed
	return us.Repo.CreateUser(u)
	//return fmt.Errorf("USER SERVICE ERROR: Register not implemented")
}

/**
* ----- UPDATE METHODS ---
*/


func (us *userService) Update(u *user.User) error {
	return us.Repo.Update(u)
}

/**
* ----- GET METHODS ---
*/

func (us *userService) GetUsers() ([]*user.User, error) {
	users, err := us.Repo.GetUsers()

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New("There is no user") 
	}

	return users, nil
}
func (us *userService) GetUserByID(id uint) (*user.User, error) {
	if id <= 0 {
		return nil, errors.New("id params is required")
	}
	user, err := us.Repo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (us *userService) GetUserByEmail(email string) (*user.User, error) {
	if email == "" {
		return nil, errors.New("email params is required")
	}
	if err := validateEmail(email); err != nil {
		return nil, err 
	}

	user, err := us.Repo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil

}
/**
* -- Other
*/

func (us *userService) HashPassword(password string) (string, error) {
	pswdAndPepper := password + us.pepper
	hashed, err := bcrypt.GenerateFromPassword([]byte(pswdAndPepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}


func (us *userService) ComparePassword(inputpswd string, dbpswd string) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(dbpswd),
		[]byte(inputpswd+us.pepper),
	)
}
func validateEmail(email string) error {
	 emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	  if !emailRegex.MatchString(email) {
			return errors.New("invalid email param entered.")
		}

		return nil
}


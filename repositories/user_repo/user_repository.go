package user_repo

import (
	"github.com/amar-jay/go-api-boilerplate/domain/user"
	"gorm.io/gorm"
)

// Repository interface
type Repo interface {
  GetUserByID(id uint) (*user.User, error)
  GetUsers() ([]*user.User, error) 
  GetUserByEmail(email string) (*user.User, error)
  CreateUser(user *user.User) error
  Update(user *user.User) error
}

type userRepo struct {
  db *gorm.DB
}

// New user repo instance
func NewUserRepo(db *gorm.DB) Repo {
  return &userRepo{
    db:db,
  }
}

func (repo *userRepo) GetUsers() ([]*user.User, error) {
  var users []*user.User
  if err:= repo.db.Find(&users).Error; err != nil {
    return nil, err
  }

  return users, nil
}
// get first user by id
func (repo *userRepo) GetUserByID(id uint) (*user.User, error) {
  var user user.User
  if err := repo.db.First(&user, id).Error; err != nil {
    return nil, err
  }

  return &user, nil
}

// Get first user by Email
func (repo *userRepo) GetUserByEmail(email string) (*user.User, error) {
  var user user.User

  if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
    return nil, err
  }

  return &user, nil 
}

// create user in db
func (repo *userRepo) CreateUser(user *user.User) error {
  return repo.db.Create(user).Error
}

// change user info
func (repo *userRepo) Update(user *user.User) error {
  return repo.db.Save(user).Error
}


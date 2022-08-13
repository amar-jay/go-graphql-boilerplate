package password_reset

import (
	pswd "github.com/amar-jay/go-api-boilerplate/domain/passwordreset"
	"gorm.io/gorm"
)

// PasswordReset Repo interface
type Repo interface {
  GetOneByToken(token string) (*pswd.PasswordReset, error)
  Create(pswd *pswd.PasswordReset) error
  Delete(id uint) error
}

type pswdRepo struct {
  db *gorm.DB
}

// create an instance of PasswordResetRepo
func CreatePasswordReserRepo(db *gorm.DB) Repo {
  return &pswdRepo{}
}

// regenerate token
func (repo *pswdRepo) GetOneByToken(token string) (*pswd.PasswordReset, error){
  var newPswd pswd.PasswordReset
  if err := repo.db.Where("token = ?", token).First(&newPswd).Error; err != nil {
     return nil, err 
}

  return &newPswd, nil
}

// Create a new password token
func (repo *pswdRepo) Create(pswd *pswd.PasswordReset) error{
  return repo.db.Create(pswd).Error
}


// Delete a token
func (repo *pswdRepo) Delete(id uint) error {
  object := pswd.PasswordReset {
    Model: gorm.Model{ID: id},
  }

  return repo.db.Delete(&object).Error
}

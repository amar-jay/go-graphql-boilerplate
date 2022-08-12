package authservice

import (
	"testing"
	"time"

	"github.com/amar-jay/go-api-boilerplate/domain/user"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
  t.Run("generating Token", func(t *testing.T){ 
  u := &user.User{
    gorm.Model{
      ID: uint(1),
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
      DeletedAt: nil,
    },
    "",
    "",
    "me@themanan.me",
    "",
    "",
    true,
  }

  svc := NewAuthService("secret")

  token, err := svc.IssueToken(*u)
  assert.Nil(t, err)
  assert.IsType(t, "string", token)

})
}

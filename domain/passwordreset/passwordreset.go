package password_reset

import "gorm.io/gorm"

// PasswordReset domain
type PasswordReset struct {
  gorm.Model
  UserID uint  `gorm:"not null"`
  Token  string `gorm:"not null,unique_index"`
}

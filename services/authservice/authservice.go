package authservice

import (
	"time"
	
	"gopkg.in/dgrijalva/jwt-go.v3"
	"github.com/amar-jay/go-api-boilerplate/domain/user"
)

// AuthService interface
type AuthService interface {
	// TODO: change userid to user.User
	IssueToken(userID int) (string, error)
	ParseToken(token string) (*Claims, error)
}

// Private authService struct
type authService struct {
	jwtSecret string
}

// Claims struct represents the claims in a JWT
type Claims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.StandardClaims

}

func NewAuthService(jwtSecret string) *authService {
	// TODO: remember to cross check return type
	return &authService{
		jwtSecret: jwtSecret,
	}
}


package authservice

import (
	"time"
	
	"gopkg.in/dgrijalva/jwt-go.v3"
	"github.com/amar-jay/go-api-boilerplate/domain/user"
	"github.com/amar-jay/go-api-boilerplate/middleware"
)

// AuthService interface
type AuthService interface {
	// TODO: change userid to user.User
	IssueToken(userID user.User) (string, error)
	ParseToken(token string) (*middleware.Claim, error)
}

// Private authService struct
type authService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) AuthService {
	// TODO: remember to cross check return type
	return &authService{
		jwtSecret: jwtSecret,
	}
}

// Generate Token for auth
func (auth *authService) IssueToken(u user.User) (string, error) {
	currTime := time.Now()
	expireTime := currTime.Add(24 * time.Hour) // after 24 hours

	claims := middleware.Claim{
		u.Email,
		int(u.ID),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "Undefined Issuer",
		},
}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
	return tokenClaims.SignedString([]byte(auth.jwtSecret)) 
}


// parse token
func (auth *authService) ParseToken(token string) (*middleware.Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(
	token,
	&middleware.Claim{},
	func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.jwtSecret), nil
	},
)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*middleware.Claim)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}


	return nil, err
}

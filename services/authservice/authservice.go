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

// Generate Token for auth
func (auth *authService) IssueToken(u user.User) (string, error) {
	currTime := time.Now()
	expireTime := currTime.Add(24 * time.Hour) // after 24 hours

	claims := Claims{
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
func (auth *authService) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
	token,
	&Claims{},
	func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.jwtSecret), nil
	},
)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}


	return nil, err
}

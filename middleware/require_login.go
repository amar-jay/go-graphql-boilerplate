package middleware

import (
	"strings"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

// Claims struct represents the claims in a JWT
type Claim struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.StandardClaims

}


// remove Bearer from "Authorization " token
func stripBearer(token string) (string, error) {
  if len(token) > 6 && strings.ToLower(token[0:7]) == "bearer " {
    return token[7:], nil
  }

  return token, nil
}

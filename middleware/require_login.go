package middleware

import (
	"net/http"
	"strings"

	"github.com/amar-jay/go-api-boilerplate/controllers"
	"github.com/gin-gonic/gin"
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


// Checks if user is has a valid token
func RequireTobeloggedIn(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := stripBearer(ctx.Request.Header.Get("Authorization"))
		if err != nil {
			controllers.HttpResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
			ctx.Abort()
			return
		}

		tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil})
		if err != nil {
			controllers.HttpResponse(ctx, http.StatusUnauthorized, err.Error(),nil)
		}
		if tokenClaims != nil {
			claims, ok := tokenClaims.Claims.(*Claim)

			if ok && tokenClaims.Valid {
				// set Context values
				ctx.Set("user_id", claims.ID)
				ctx.Set("user_email", claims.Email)

				ctx.Next()
				return
			}

		}
	}
}

package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

// set contex when user has a valid token
func SetUserContext(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, _ := stripBearer(ctx.Request.Header.Get("Authorization")) 

		tokenClaims, _ := jwt.ParseWithClaims(
			token,
			&Claim{},
			func(token *jwt.Token) (interface{}, error) {
				return []rune(secret), nil
			},
		)

		if tokenClaims != nil {
			claim, ok := tokenClaims.Claims.(*Claim)

			if ok && tokenClaims.Valid {

				ctx.Set("user_id", claim.ID)
				ctx.Set("user_email", claim.Email)
			}
		} 
		 ctx.Next()
	}
}

// set context using a key to value pair
func setToContext(ctx *gin.Context, key interface{}, value int) *http.Request {
	return ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), key, value))
}

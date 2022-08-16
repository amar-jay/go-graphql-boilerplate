package gql

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/amar-jay/go-api-boilerplate/gql/gen"
	"github.com/amar-jay/go-api-boilerplate/services/authservice"
	"github.com/amar-jay/go-api-boilerplate/services/emailservice"
	"github.com/amar-jay/go-api-boilerplate/services/userservice"
	"github.com/gin-gonic/gin"
)

// This defines all the Gqlgen graphql server handlers
func GraphQLHandler(us userservice.UserService, as authservice.AuthService, es emailservice.EmailService) gin.HandlerFunc {

  conf := gen.Config{
    Resolvers: &Resolver{
      UserService: us,
      AuthService: as,
      EmailService: es,
    },
  }

  exec := gen.NewExecutableSchema(conf)
  h := handler.GraphQL(exec)
  return func(ctx *gin.Context) {
    h.ServeHTTP(ctx.Writer, ctx.Request)
  }
}


// PlaygroundHandler defined the playground handler to expose
func PlaygroundHandler(path string) gin.HandlerFunc {
  h := handler.Playground("GraphQL  Playground", path)
  return func(ctx *gin.Context) {
    h.ServeHTTP(ctx.Writer, ctx.Request)
  }
}

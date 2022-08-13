package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/Massad/gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/amar-jay/go-api-boilerplate/config"
	"github.com/amar-jay/go-api-boilerplate/controllers"
	"github.com/amar-jay/go-api-boilerplate/domain/user"
	"github.com/amar-jay/go-api-boilerplate/gql"
	"github.com/amar-jay/go-api-boilerplate/middleware"
	"github.com/amar-jay/go-api-boilerplate/services/authservice"
	"github.com/amar-jay/go-api-boilerplate/services/emailservice"
	"github.com/amar-jay/go-api-boilerplate/services/userservice"
)

const defaultPort = "8080"
var (
	router = gin.Default()
)
func main() {
	fmt.Println("Starting server...")
	router.SetTrustedProxies([]string{"192.168.1.2"})

	// swagger url - http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))



	// load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading env files. $PORT must be set")
	}
	config := config.GetConfig()

	db, err := gorm.Open(
		config.Postgres.GetConnectionInfo(),
		config.Postgres.Config(), 
	)

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&user.User{})
//	defer db.Close()
	fmt.Println("Database migrated successfully")

	router.GET("/", func(c *gin.Context) {
		// If the client is 192.168.1.2, use the X-Forwarded-For
		// header to deduce the original client IP from the trust-
		// worthy parts of that header.
		// Otherwise, simply return the direct client IP
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{"Amar": "Jay", "clientIP": c.ClientIP()})
	})

	// Testing the database
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	/**
	*  ----- Services -----
	*/

	userService  := userservice.NewUserService("pepper")
	authService := authservice.NewAuthService(config.JWTSecret)
	emailService :=  emailservice.NewEmailService()

	/**
	* ----- Controllers -----
	*/


	userController := controllers.NewUserController(userService, authService, emailService)

	/**
	*  ----- Routing -----
	*/

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	playground :=handler.Playground("GraphQL playground", "/query")
	router.GET("/graphql", func (c *gin.Context) { playground.ServeHTTP(c.Writer, c.Request) })
	router.POST("/query", func (c *gin.Context) { 
		middleware.SetUserContext(config.JWTSecret)
		gql.GraphQLHandler(userService, authService, emailService)
		playground.ServeHTTP(c.Writer, c.Request) })
	// http.Handle("/query", srv)

	api := router.Group("/api")

	api.POST("/register", userController.Register)
	api.POST("/login",  userController.Login)
	api.POST("/forgot-password", userController.ForgotPassword)
	api.POST("/update-password", userController.ResetPassword)

	user := api.Group("/user")

	user.GET("/:id", userController.GetUserById)

	// TODO: create accounts and profiles

	// log.Printf("connect to http://loc alhost:%s/ for GraphQL playground", port)
	port := fmt.Sprintf(":%d", config.Port)
	router.Run(port)
}

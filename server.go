package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/amar-jay/go-api-boilerplate/common/hmachash"
	"github.com/amar-jay/go-api-boilerplate/common/randomstring"
	"github.com/amar-jay/go-api-boilerplate/config"
	"github.com/amar-jay/go-api-boilerplate/controllers"
	"github.com/amar-jay/go-api-boilerplate/domain/user"
	"github.com/amar-jay/go-api-boilerplate/gql"
	"github.com/amar-jay/go-api-boilerplate/middleware"
	"github.com/amar-jay/go-api-boilerplate/repositories/password_reset"
	"github.com/amar-jay/go-api-boilerplate/repositories/user_repo"
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
	router.SetTrustedProxies([]string{"192.168.1.2", "::1"})

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
	userrepo := user_repo.NewUserRepo(db)
	pswdrepo := password_reset.CreatePasswordReserRepo(db)
	randomstr := randomstring.CreateRandomString()
	hash := hmachash.NewHMAC("dumb ass")
	userService := userservice.NewUserService(userrepo, pswdrepo, randomstr, hash, config.Pepper)
	authService := authservice.NewAuthService(config.JWTSecret)
	emailService := emailservice.NewEmailService()

	/**
	* ----- Controllers -----
	 */

	userController := controllers.NewUserController(userService, authService, emailService)

	/**
	*  ----- Routing -----
	 */

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//playground := handler.Playground("GraphQL playground", "/query")
	router.GET("/graphql",  gql.PlaygroundHandler("/query"))
	router.POST("/query", func(c *gin.Context) {
		middleware.SetUserContext(config.JWTSecret)
		gql.GraphQLHandler(userService, authService, emailService)
	})
	// http.Handle("/query", srv)

	auth := router.Group("/auth")

	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	auth.POST("/forgot-password", userController.ForgotPassword)
	auth.POST("/update-password", userController.ResetPassword)

	user := router.Group("/users")

	user.GET("/", userController.GetUsers)
	user.GET("/:id", userController.GetUserByID)

	//  accounts and profiles
	account := router.Group("/account")
	account.Use(middleware.RequireTobeloggedIn(config.JWTSecret))
	{
		account.GET("/profile", userController.GetProfile)
		account.PUT("/profile", userController.Update)
	}

	// Run server
	log.Printf("Running on http://localhost:%d/ ", config.Port)
	port := fmt.Sprintf(":%d", config.Port)
	router.Run(port)
}

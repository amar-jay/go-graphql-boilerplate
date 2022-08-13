package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/amar-jay/go-api-boilerplate/config"
	"github.com/amar-jay/go-api-boilerplate/domain/user"
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

	 //srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

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
		// log.Fatal("Error in connecting to database")
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
	})
	// Testing the database
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	playground :=handler.Playground("GraphQL playground", "/query")
	router.GET("/graphql", func (c *gin.Context) { playground.ServeHTTP(c.Writer, c.Request) })
	// http.Handle("/query", srv)

	// log.Printf("connect to http://loc alhost:%s/ for GraphQL playground", port)
	port := fmt.Sprintf(":%d", config.Port)
	router.Run(port)
}

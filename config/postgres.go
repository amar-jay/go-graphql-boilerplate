package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


// PostgresConfig is the config for postgres
type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
}
//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// Dialect returns the dialect for postgres
func (c *PostgresConfig) Config() *gorm.Config {
	return &gorm.Config{}
}

// GetConnectionInfo returns the connection info for postgres
func (c *PostgresConfig) GetConnectionInfo() gorm.Dialector {
	fmt.Println(c.Host, c.Port, c.User, c.Password, c.Database)
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Database)
	return postgres.Open(dns) 
}

// GetPostgresConfig returns the postgres config
func GetPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     getPort("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}
}

func getPort(input string) int {
	port, err := strconv.Atoi(os.Getenv(input))
	if err != nil {
		fmt.Println("Error in parsing POSTGRES_PORT. Setting to 5432")
		panic(err)
	}
	if port < 1000 {
		panic(fmt.Errorf("Invalid port number: %d", port))
	}

	return port
}

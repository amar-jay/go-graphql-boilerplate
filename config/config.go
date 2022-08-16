package config

import (
	"os"
)
const (
  // AppName is the name of the app
  appName = "go-api-boilerplate"
  production = "production"
)

type Config struct {
  Pepper    string        `env:"PEPPER"`
  Env       string        `env:"ENV"`
  FromEmail string        `env:"EMAIL_FROM"`
  Port      int           `env:"PORT"`
  JWTSecret string        `env:"JWT_SECRET"`
  Mailgun   MailgunConfig `json:"mailgun"`
  Postgres  PostgresConfig `json:"postgres"`
}

// Check if it is in production
func (c Config) isProduction() bool {
  return c.Env == production
}

func GetConfig() Config {
  return Config{
    Pepper:  os.Getenv("PEPPER"),
    Env: os.Getenv("ENV"),
    Mailgun: GetMailgunConfig(),
    Postgres: GetPostgresConfig(),
    FromEmail: os.Getenv("EMAIL_FROM"),
    Port: getPort("PORT"),
    JWTSecret: os.Getenv("JWT_SECRET"),
  }
}

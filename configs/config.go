package config

type Config {
  Env string `env:"ENV"`
  Mailgun MailgunConfig `json:"mailgun"`
  FromEmail string `env:"EMAIL_FROM"`
}

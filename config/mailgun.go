package config

type MailgunConfig struct {
	APIKey    string `env:"MAILGUN_API_KEY"`
	PublicKey string `env:"MAILGUN_PUBLIC_KEY"`
	Domain    string `env:"MAILGUN_DOMAIN"`
}

func GetMailgunConfig() MailgunConfig {
	return MailgunConfig{
		APIKey:    "",
		PublicKey: "",
		Domain:    "",
	}
}

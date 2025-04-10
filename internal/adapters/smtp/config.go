package smtp

import (
	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Host      string
	Port      int
	User      string
	Pass      string
	FromEmail string
}

func LoadFromEnv() *Config {
	return &Config{
		Host:      env.GetEnv[string]("EZEX_NOTIFICATION_SMTP_HOST"),
		Port:      env.GetEnv[int]("EZEX_NOTIFICATION_SMTP_PORT", env.WithDefault("587")),
		User:      env.GetEnv[string]("EZEX_NOTIFICATION_SMTP_USER"),
		Pass:      env.GetEnv[string]("EZEX_NOTIFICATION_SMTP_PASS"),
		FromEmail: env.GetEnv[string]("EZEX_NOTIFICATION_SMTP_FROM_EMAIL"),
	}
}

func (*Config) BasicCheck() error {
	// Add validation if needed
	return nil
}

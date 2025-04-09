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
		Host:      env.GetEnv[string]("SMTP_HOST"),
		Port:      env.GetEnv[int]("SMTP_PORT", env.WithDefault("587")),
		User:      env.GetEnv[string]("SMTP_USER"),
		Pass:      env.GetEnv[string]("SMTP_PASS"),
		FromEmail: env.GetEnv[string]("SMTP_FROM_EMAIL"),
	}
}

func (*Config) BasicCheck() error {
	// Add validation if needed
	return nil
}

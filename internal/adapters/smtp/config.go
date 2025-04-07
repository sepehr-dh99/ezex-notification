package smtp

import (
	"os"
	"strconv"
)

type Config struct {
	Host      string
	Port      int
	User      string
	Pass      string
	FromEmail string
}

func DefaultConfig() *Config {
	return &Config{
		Host:      "smtp.example.com",
		Port:      587,
		User:      "smtp_user",
		Pass:      "smtp_password",
		FromEmail: "no-reply@example.com",
	}
}

func LoadFromEnv() *Config {
	cfg := DefaultConfig()

	if v := os.Getenv("SMTP_HOST"); v != "" {
		cfg.Host = v
	}

	if v := os.Getenv("SMTP_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			cfg.Port = port
		}
	}

	if v := os.Getenv("SMTP_USER"); v != "" {
		cfg.User = v
	}

	if v := os.Getenv("SMTP_PASS"); v != "" {
		cfg.Pass = v
	}

	if v := os.Getenv("SMTP_FROM_EMAIL"); v != "" {
		cfg.FromEmail = v
	}

	return cfg
}

func (*Config) BasicCheck() error {
	// Add validation if needed
	return nil
}

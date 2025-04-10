package grpc

import (
	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Port string
}

func LoadFromEnv() *Config {
	return &Config{
		Port: env.GetEnv[string]("EZEX_NOTIFICATION_GRPC_PORT", env.WithDefault("50051")),
	}
}

func (*Config) BasicCheck() error {
	// Add validation if needed
	return nil
}

package grpc

import (
	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Address string
}

func LoadFromEnv() *Config {
	return &Config{
		Address: env.GetEnv[string]("EZEX_NOTIFICATION_GRPC_ADDRESS", env.WithDefault("0.0.0.0:50051")),
	}
}

func (*Config) BasicCheck() error {
	// Add validation if needed
	return nil
}

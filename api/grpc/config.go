package grpc

import (
	"errors"

	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Port string
}

func LoadFromEnv() *Config {
	return &Config{
		Port: env.GetEnv[string]("GRPC_PORT", env.WithDefault("50051")),
	}
}

func (c *Config) BasicCheck() error {
	if c.Port == "" {
		return errors.New("config: gRPC port is not set")
	}

	return nil
}

package grpc

import (
	"errors"
	"os"
)

type Config struct {
	Port string
}

func DefaultConfig() *Config {
	return &Config{
		Port: "50051",
	}
}

func LoadFromEnv() *Config {
	cfg := DefaultConfig()

	if v := os.Getenv("GRPC_PORT"); v != "" {
		cfg.Port = v
	}

	return cfg
}

func (c *Config) BasicCheck() error {
	if c.Port == "" {
		return errors.New("config: gRPC port is not set")
	}

	return nil
}

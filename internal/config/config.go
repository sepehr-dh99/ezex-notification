// Package config defines project configurations
package config

import (
	"github.com/ezex-io/ezex-notification/api/grpc"
	"github.com/ezex-io/ezex-notification/internal/adapters/smtp"
)

type Config struct {
	SMTP *smtp.Config
	GRPC *grpc.Config
}

func Load() *Config {
	config := &Config{
		SMTP: smtp.LoadFromEnv(),
		GRPC: grpc.LoadFromEnv(),
	}

	return config
}

// BasicCheck checks the necessary config checking from each module.
func (cfg *Config) BasicCheck() error {
	if err := cfg.SMTP.BasicCheck(); err != nil {
		return err
	}

	if err := cfg.GRPC.BasicCheck(); err != nil {
		return err
	}

	return nil
}

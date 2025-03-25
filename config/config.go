// Package config defines project global variables
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config schema type.
type Config struct {
	GRPCPort  string
	SMTPHost  string
	SMTPPort  int
	SMTPUser  string
	SMTPPass  string
	OTPLength int
	FromEmail string
}

// Load retieves project configs from environment variables.
func Load() (*Config, error) {
	// Load from the current working directory
	err := godotenv.Load("config/.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load comfigs: %w", err)
	}

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	otpLength, _ := strconv.Atoi(os.Getenv("OTP_LENGTH"))

	return &Config{
		GRPCPort:  os.Getenv("GRPC_PORT"),
		SMTPHost:  os.Getenv("SMTP_HOST"),
		SMTPPort:  port,
		SMTPUser:  os.Getenv("SMTP_USER"),
		SMTPPass:  os.Getenv("SMTP_PASS"),
		OTPLength: otpLength,
		FromEmail: os.Getenv("FROM_EMAIL"),
	}, nil
}

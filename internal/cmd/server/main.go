// Package main defines project entry
package main

import (
	"flag"

	grpcserver "github.com/ezex-io/ezex-notification/api/grpc"
	"github.com/ezex-io/ezex-notification/internal/adapters/smtp"
	"github.com/ezex-io/ezex-notification/internal/config"
	"github.com/ezex-io/ezex-notification/internal/interactors"
	"github.com/ezex-io/gopkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	envFile := flag.String("env", ".env", "Path to environment file")
	flag.Parse()

	logging := logger.NewSlog(nil)

	// Load the specified env file
	if err := godotenv.Load(*envFile); err != nil {
		logging.Warn("Failed to load env file '%s': %v. Continuing with system environment...", *envFile, err)
	} else {
		logging.Debug("Loaded environment variables from '%s'", *envFile)
	}

	// Load config from environment
	cfg := config.Load()
	if err := cfg.BasicCheck(); err != nil {
		logging.Fatal("failed to load config: %v", err)
	}

	emailSender := smtp.NewSMTPAdapter(cfg.SMTP)
	emailWorker := interactors.NewEmailWorker(emailSender)

	notificationService := grpcserver.NewNotificationService(emailWorker)

	server, err := grpcserver.NewServer(notificationService, grpcserver.Config{
		Port: cfg.GRPC.Port,
	})
	if err != nil {
		logging.Fatal("failed to create gRPC server: %v", err)
	}

	go server.Start()
	logging.Debug("Starting gRPC server on port %s", cfg.GRPC.Port)

	if err := <-server.Notify(); err != nil {
		logging.Fatal("gRPC server error: %v", err)
	}
}

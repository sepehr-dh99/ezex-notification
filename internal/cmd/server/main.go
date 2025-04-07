// Package main defines project entry
package main

import (
	"flag"
	"log"

	grpcserver "github.com/ezex-io/ezex-notification/api/grpc"
	"github.com/ezex-io/ezex-notification/internal/adapters/smtp"
	"github.com/ezex-io/ezex-notification/internal/config"
	"github.com/ezex-io/ezex-notification/internal/interactors"
	"github.com/joho/godotenv"
)

func main() {
	envFile := flag.String("env", ".env", "Path to environment file")
	flag.Parse()

	// Load the specified env file
	if err := godotenv.Load(*envFile); err != nil {
		log.Printf("Failed to load env file '%s': %v. Continuing with system environment...", *envFile, err)
	} else {
		log.Printf("Loaded environment variables from '%s'", *envFile)
	}

	// Load config from environment
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	emailSender := smtp.NewSMTPAdapter(cfg.SMTP)
	emailWorker := interactors.NewEmailWorker(emailSender)

	notificationService := grpcserver.NewNotificationService(emailWorker)

	server, err := grpcserver.NewServer(notificationService, grpcserver.Config{
		Port: cfg.GRPC.Port,
	})
	if err != nil {
		log.Fatalf("failed to create gRPC server: %v", err)
	}

	go server.Start()
	log.Printf("Starting gRPC server on port %s", cfg.GRPC.Port)

	if err := <-server.Notify(); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}

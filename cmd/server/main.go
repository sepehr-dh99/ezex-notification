package main

import (
	"log"
	"net"

	"github.com/ezex-io/ezex-notification/config"
	"github.com/ezex-io/ezex-notification/internal/adapters/email"
	grpcserver "github.com/ezex-io/ezex-notification/internal/adapters/grpc"
	"github.com/ezex-io/ezex-notification/internal/adapters/grpc/proto"
	"github.com/ezex-io/ezex-notification/internal/interactors"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize email adapter
	emailSender := email.NewSMTPAdapter(cfg)

	// Initialize OTP service
	otpService := interactors.NewOTPService(emailSender, cfg)

	// Create gRPC server
	server := googleGrpc.NewServer()
	proto.RegisterNotificationServiceServer(server, grpcserver.NewServer(otpService))

	reflection.Register(server)

	// Start listening
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting gRPC server on port %s", cfg.GRPCPort)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

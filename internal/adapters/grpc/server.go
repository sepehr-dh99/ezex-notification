// Package grpc defines gRPC server
package grpc

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-notification/internal/adapters/grpc/proto"
	"github.com/ezex-io/ezex-notification/internal/interactors"
)

// Server is schema type of gRPC server.
type Server struct {
	proto.UnimplementedNotificationServiceServer
	otpService *interactors.OTPService
}

// NewServer acts as constructor of gRPC server.
func NewServer(otpService *interactors.OTPService) *Server {
	return &Server{otpService: otpService}
}

// SendOTP uses GenerateAndSendOTP and return code to it caller.
func (s *Server) SendOTP(_ context.Context, req *proto.SendOTPRequest) (*proto.SendOTPResponse, error) {
	otp, err := s.otpService.GenerateAndSendOTP(req.Email)
	if err != nil {
		return &proto.SendOTPResponse{
			Otp: "",
		}, fmt.Errorf("service error: %w", err)
	}

	return &proto.SendOTPResponse{
		Otp: otp,
	}, nil
}

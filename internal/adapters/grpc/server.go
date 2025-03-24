package grpc

import (
	"context"

	"github.com/ezex-io/ezex-notification/internal/adapters/grpc/proto"
	"github.com/ezex-io/ezex-notification/internal/interactors"
)

type Server struct {
	proto.UnimplementedNotificationServiceServer
	otpService *interactors.OTPService
}

func NewServer(otpService *interactors.OTPService) *Server {
	return &Server{otpService: otpService}
}

func (s *Server) SendOTP(ctx context.Context, req *proto.SendOTPRequest) (*proto.SendOTPResponse, error) {
	otp, err := s.otpService.GenerateAndSendOTP(req.Email)
	if err != nil {
		return &proto.SendOTPResponse{
			Otp:     "",
			Success: false,
		}, err
	}

	return &proto.SendOTPResponse{
		Otp:     otp,
		Success: true,
	}, nil
}

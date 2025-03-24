package interactors

import (
	"math/rand"
	"time"

	"github.com/ezex-io/ezex-notification/config"
	"github.com/ezex-io/ezex-notification/internal/ports"
)

type OTPService struct {
	emailSender ports.EmailSender
	cfg         *config.Config
}

func NewOTPService(emailSender ports.EmailSender, cfg *config.Config) *OTPService {
	return &OTPService{
		emailSender: emailSender,
		cfg:         cfg,
	}
}

func (s *OTPService) GenerateAndSendOTP(email string) (string, error) {
	otp := generateOTP(s.cfg.OTPLength)

	if err := s.emailSender.SendOTPEmail(email, otp); err != nil {
		return "", err
	}

	return otp, nil
}

func generateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	otp := make([]byte, length)

	for i := range otp {
		otp[i] = digits[rand.Intn(len(digits))]
	}
	return string(otp)
}

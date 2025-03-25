// Package interactors define business logics
package interactors

import (
	"crypto/rand"
	"fmt"

	"github.com/ezex-io/ezex-notification/config"
	"github.com/ezex-io/ezex-notification/internal/ports"
)

// OTPService is otp service schema type.
type OTPService struct {
	emailSender ports.EmailSender
	cfg         *config.Config
}

// NewOTPService acts as constructor of otp service.
func NewOTPService(emailSender ports.EmailSender, cfg *config.Config) *OTPService {
	return &OTPService{
		emailSender: emailSender,
		cfg:         cfg,
	}
}

// GenerateAndSendOTP call's generateOTP and pass the generated otp to SendOTPEmail.
func (s *OTPService) GenerateAndSendOTP(email string) (string, error) {
	otp := generateOTP(s.cfg.OTPLength)

	if err := s.emailSender.SendOTPEmail(email, otp); err != nil {
		return "", fmt.Errorf("failed to send OTP email to %s: %w", email, err)
	}

	return otp, nil
}

// generateOTP generates 6 digits random OTP code using 0 to 9 numbers.
func generateOTP(length int) string {
	digits := "0123456789"
	otp := make([]byte, length)

	_, err := rand.Read(otp)
	if err != nil {
		panic("failed to generate random bytes")
	}

	for i := range otp {
		otp[i] = digits[otp[i]%byte(len(digits))]
	}

	return string(otp)
}

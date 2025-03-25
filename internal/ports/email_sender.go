// Package ports defines interfaces (ports)
package ports

// EmailSender defines an interface for sending OTP emails.
type EmailSender interface {
	SendOTPEmail(email, otp string) error
}

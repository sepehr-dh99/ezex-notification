package ports

type EmailSender interface {
	SendOTPEmail(email, otp string) error
}

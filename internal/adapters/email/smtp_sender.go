package email

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/ezex-io/ezex-notification/config"
	"github.com/ezex-io/ezex-notification/internal/ports"
	"gopkg.in/gomail.v2"
)

type SMTPAdapter struct {
	config *config.Config
}

func NewSMTPAdapter(cfg *config.Config) ports.EmailSender {
	return &SMTPAdapter{config: cfg}
}

func (a *SMTPAdapter) SendOTPEmail(to, otp string) error {
	tmplPath := filepath.Join("assets", "template", "email", "otp.md")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, map[string]string{"OTP": otp}); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", a.config.FromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your Verification Code")
	m.SetBody("text/plain", body.String())

	d := gomail.NewDialer(
		a.config.SMTPHost,
		a.config.SMTPPort,
		a.config.SMTPUser,
		a.config.SMTPPass,
	)

	return d.DialAndSend(m)
}

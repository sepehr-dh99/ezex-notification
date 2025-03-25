// Package email defines email adaptor
package email

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"

	embed "github.com/ezex-io/ezex-notification"
	"github.com/ezex-io/ezex-notification/config"
	"github.com/ezex-io/ezex-notification/internal/ports"
	"gopkg.in/gomail.v2"
)

// SMTPAdapter schema for email smtp adapter.
type SMTPAdapter struct {
	config        *config.Config
	templateCache map[string]*template.Template
	templateMutex sync.RWMutex
}

// NewSMTPAdapter act as smtp sender constructor.
func NewSMTPAdapter(cfg *config.Config) ports.EmailSender {
	return &SMTPAdapter{
		config:        cfg,
		templateCache: make(map[string]*template.Template),
	}
}

// SendOTPEmail uses markdown template and otp code to send an email.
func (a *SMTPAdapter) SendOTPEmail(sendTo, otp string) error {
	tmpl, err := a.getTemplate("otp.md")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, map[string]string{"OTP": otp}); err != nil {
		return fmt.Errorf("failed to apply otp code to template: %w", err)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", a.config.FromEmail)
	message.SetHeader("To", sendTo)
	message.SetHeader("Subject", "Your Verification Code")
	message.SetBody("text/plain", body.String())

	dialer := gomail.NewDialer(
		a.config.SMTPHost,
		a.config.SMTPPort,
		a.config.SMTPUser,
		a.config.SMTPPass,
	)

	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send OTP email to %s: %w", sendTo, err)
	}

	return nil
}

func (a *SMTPAdapter) getTemplate(filename string) (*template.Template, error) {
	// Check cache first
	a.templateMutex.RLock()
	cachedTmpl, exists := a.templateCache[filename]
	a.templateMutex.RUnlock()
	if exists {
		return cachedTmpl, nil
	}

	// Load from embedded FS
	content, err := embed.TemplateFS.ReadFile("assets/template/email/" + filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read email template %s: %w", filename, err)
	}

	// Parse template
	tmpl, err := template.New(filename).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse email template to string: %w", err)
	}

	// Store in cache
	a.templateMutex.Lock()
	a.templateCache[filename] = tmpl
	a.templateMutex.Unlock()

	return tmpl, nil
}

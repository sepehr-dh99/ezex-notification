package email

import (
	"bytes"
	"sync"
	"text/template"

	embed "github.com/ezex-io/ezex-notification"
	"github.com/ezex-io/ezex-notification/config"
	"github.com/ezex-io/ezex-notification/internal/ports"
	"gopkg.in/gomail.v2"
)

type SMTPAdapter struct {
	config        *config.Config
	templateCache map[string]*template.Template
	templateMutex sync.RWMutex
}

func NewSMTPAdapter(cfg *config.Config) ports.EmailSender {
	return &SMTPAdapter{
		config:        cfg,
		templateCache: make(map[string]*template.Template),
	}
}

func (a *SMTPAdapter) SendOTPEmail(to, otp string) error {
	tmpl, err := a.getTemplate("otp.md")
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
		return nil, err
	}

	// Parse template
	tmpl, err := template.New(filename).Parse(string(content))
	if err != nil {
		return nil, err
	}

	// Store in cache
	a.templateMutex.Lock()
	a.templateCache[filename] = tmpl
	a.templateMutex.Unlock()

	return tmpl, nil
}

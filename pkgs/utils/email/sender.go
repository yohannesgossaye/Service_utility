package email

import (
	"fmt"
	"net/smtp"
	"os"
)

type Sender interface {
	Send(to, subject, body string) error
}

type SMTPSender struct {
	auth smtp.Auth
	host string
	port string
	from string
}

// Build SMTP sender from environment
func NewSMTPSender() *SMTPSender {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	// Validate env vars
	if host == "" || port == "" || user == "" || pass == "" {
		fmt.Println("⚠️ Missing SMTP configuration in environment variables")
		return nil
	}

	auth := smtp.PlainAuth("", user, pass, host)

	return &SMTPSender{
		auth: auth,
		host: host,
		port: port,
		from: user,
	}
}

// Send actual email
func (s *SMTPSender) Send(to, subject, body string) error {
	addr := s.host + ":" + s.port
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))
	return smtp.SendMail(addr, s.auth, s.from, []string{to}, msg)
}

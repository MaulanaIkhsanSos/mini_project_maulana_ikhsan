// email_util.go

package util

import (
	"net/smtp"
)

// EmailSender adalah interface yang digunakan untuk mengirim email.
type EmailSender interface {
	SendEmail(subject, body string, to []string) error
}

// SimpleEmailSender adalah implementasi sederhana dari EmailSender interface untuk mengirim email menggunakan SMTP.
type SimpleEmailSender struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

// NewSimpleEmailSender adalah fungsi untuk membuat instansiasi SimpleEmailSender.
func NewSimpleEmailSender(smtpServer, smtpPort, smtpUsername, smtpPassword string) *SimpleEmailSender {
	return &SimpleEmailSender{
		SMTPServer:   smtpServer,
		SMTPPort:     smtpPort,
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
	}
}

// SendEmail digunakan untuk mengirim email.
func (s *SimpleEmailSender) SendEmail(subject, body string, to []string) error {
	auth := smtp.PlainAuth("", s.SMTPUsername, s.SMTPPassword, s.SMTPServer)

	msg := "From: " + s.SMTPUsername + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	return smtp.SendMail(s.SMTPServer+":"+s.SMTPPort, auth, s.SMTPUsername, to, []byte(msg))
}

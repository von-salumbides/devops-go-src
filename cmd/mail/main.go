package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/von-salumbides/devops-go-src/utils/logger"
	"go.uber.org/zap"
)

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func main() {
	logger.InitLogger()
	// Sender data.
	from := os.Getenv("MAIL_FROM")
	// Receiver email address.
	to := []string{os.Getenv("MAIL_TO")}
	password := os.Getenv("MAIL_PASSWORD")
	// Body - message to provide
	body := "This is a test email message."
	// Subject
	subject := os.Getenv("MAIL_SUBJECT")
	// smtp server configuration.
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")

	request := Mail{
		From:    from,
		To:      to,
		Body:    body,
		Subject: subject,
	}
	msg := BuildMessage(request)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		zap.L().Error("Failed to send email", zap.Any("error", err.Error()))
		return
	}
	zap.L().Info("Email Sent Successfully!")
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

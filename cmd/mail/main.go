package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/von-salumbides/devops-go-src/config"
	"github.com/von-salumbides/devops-go-src/utils/logger"
	"go.uber.org/zap"
)

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

var yaml []byte

func main() {
	logger.InitLogger()
	config, err := config.ConfigSetup(os.Getenv("ENVIRONMENT"), "mail")
	if err != nil {
		zap.L().Error("Error loading config file", zap.Any("error", err.Error()))
		os.Exit(1)
	}
	// Mail config
	mailToSlice := strings.Split(config.GetString("mail.TO"), ",")
	to := mailToSlice
	from := config.GetString("mail.FROM")
	password := config.GetString("mail.PASSWORD")
	subject := config.GetString("mail.SUBJECT")
	smtpHost := config.GetString("mail.HOST")
	smtpPort := config.GetString("mail.PORT")
	body := "This is a test email message."
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
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
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

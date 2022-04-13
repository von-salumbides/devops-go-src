package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"strings"

	"github.com/von-salumbides/devops-go-src/configs"
	"github.com/von-salumbides/devops-go-src/internal/logger"
)

type Mail struct {
	From     string   `json:"from"`
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	Password string   `json:"password"`
	Host     string   `json:"host"`
	Port     string   `json:"port"`
	Message  string   `json:"message"`
	Body     bytes.Buffer
}

type DataTemplate struct {
	Message     string
	Environment string
}

func main() {
	logger.InitLogger()
	config, err := configs.ConfigSetup(os.Getenv("ENVIRONMENT"), "MAIL")
	if err != nil {
		logger.ERROR("Error loading config file", err.Error())
		os.Exit(1)
	}
	// Mail config
	// TODO: Change to template
	var tmplBody bytes.Buffer
	tmpl := template.Must(template.ParseFiles("cmd/mail/mail.tmpl"))
	data := DataTemplate{
		Message:     config.GetString("mail.MESSAGE"),
		Environment: config.GetString("mail.ENVIRONMENT"),
	}
	err = tmpl.Execute(&tmplBody, data)
	if err != nil {
		logger.ERROR("Error loading template file", err.Error())
		os.Exit(1)
	}
	request := Mail{
		From:     config.GetString("mail.FROM"),
		To:       strings.Split(config.GetString("mail.TO"), ","),
		Subject:  config.GetString("mail.SUBJECT"),
		Password: config.GetString("mail.PASSWORD"),
		Host:     config.GetString("mail.HOST"),
		Port:     config.GetString("mail.PORT"),
		Body:     tmplBody,
	}
	msg := BuildMessage(request)
	// Authentication.
	auth := smtp.PlainAuth("",
		request.From,
		request.Password,
		request.Host)
	// Sending email.
	err = smtp.SendMail(request.Host+":"+request.Port,
		auth,
		request.From,
		request.To,
		[]byte(msg))
	if err != nil {
		logger.ERROR("Failed to send email", err.Error())
		return
	}
	logger.INFO("Email Sent Successfully!")
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body.String())

	return msg
}

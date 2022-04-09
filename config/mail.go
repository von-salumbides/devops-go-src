package config

type MailCOnfig struct {
	MailFrom     string `mapstructure:"MAIL_FROM"`
	MailTo       string `mapstructure:"MAIL_TO"`
	MailPassword string `mapstructure:"MAIL_PASSWORD"`
	MailSubject  string `mapstructure:"MAIL_SUBJECT"`
	MailHost     string `mapstructure:"MAIL_HOST"`
	MailsPort    string `mapstructure:"MAIL_PORT"`
}

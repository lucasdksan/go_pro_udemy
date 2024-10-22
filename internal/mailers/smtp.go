package mailers

import "gopkg.in/gomail.v2"

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type smtpMailService struct {
	from   string
	dialer *gomail.Dialer
}

func NewSMTPMailService(cfg SMTPConfig) MailService {
	dailer := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)

	return smtpMailService{from: cfg.From, dialer: dailer}
}

func (sms smtpMailService) Send(msg MailMessage) error {
	m := gomail.NewMessage()

	m.SetHeader("From", sms.from)
	m.SetHeader("To", msg.To...)
	m.SetHeader("Subject", msg.Subject)

	if msg.IsHTML {
		m.SetBody("text/html", string(msg.Body))
	} else {
		m.SetBody("text/plain", string(msg.Body))
	}

	return sms.dialer.DialAndSend(m)
}

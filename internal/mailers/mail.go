package mailers

type MailMessage struct {
	To      []string
	Subject string
	Body    []byte
	IsHTML  bool
}

type MailService interface {
	Send(msg MailMessage) error
}

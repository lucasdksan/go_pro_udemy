package mailers

type consoleMailService struct {
	from string
}

func NewConsoleMailService(from string) MailService {
	return consoleMailService{from: from}
}

func (cms consoleMailService) Send(msg MailMessage) error {
	return nil
}

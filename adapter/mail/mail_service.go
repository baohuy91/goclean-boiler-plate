package mail

import "goclean/usecase"

// A mail client that sent the mail object away
type MailGateway interface {
	SendMail(mail MailMsg) error
}

func NewMailService(mailClient MailGateway) usecase.MailService {
	return &mailServiceImpl{
		mailClient: mailClient,
	}
}

type mailServiceImpl struct {
	mailClient MailGateway
}

func (m *mailServiceImpl) SendMail(mailContent string, toUid string) error {
	toAddress, _ := getUserMailAddress(toUid)
	msg := &mailMsgImpl{
		content: mailContent,
		toList:  []string{toAddress},
	}
	err := m.mailClient.SendMail(msg)
	if err != nil {
		return err
	}

	return nil
}

func getUserMailAddress(uid string) (string, error) {
	// TODO: implement
	return "", nil
}

package mail

// Return a builder to build a mail object, and then give it to mail manager to be sent away
func NewBuilder(mailContent string, toAddress string) MailBuilder {
	return &mailBuilderImpl{
		mail: &mailImpl{
			content: mailContent,
			toList:  []string{toAddress},
		},
	}
}

type MailBuilder interface {
	Build() Mail
}

type mailBuilderImpl struct {
	mail *mailImpl
}

func (m *mailBuilderImpl) Build() Mail {
	return *m.mail
}

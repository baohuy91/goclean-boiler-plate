package controller

type mailManagerMock struct {
	sendMailFunc func() error
}

func (m *mailManagerMock) SendMail(mail Mail) error {
	if m.sendMailFunc != nil {
		return m.sendMailFunc()
	}
	return nil
}

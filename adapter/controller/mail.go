package controller

type MailManager interface {
	SendMail(mail Mail) error
}

type Mail struct {
	From         string
	ToList       []string
	CCList       []string
	Subject      string
	Content      string
	CustomArgs   map[string]string
	Categories   []string
	InReplyTo    string
	ReferenceIds []string
}

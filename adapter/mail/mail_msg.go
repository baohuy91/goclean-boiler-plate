package mail

// MailMsg struct that content message to feed to mail client to be sent away
type MailMsg interface {
	From() string
	ToList() []string
	CCList() []string
	Subject() string
	Content() string
	CustomArgs() map[string]string
	InReplyTo() string
	Categories() []string
	ReferenceIds() []string
}

type mailMsgImpl struct {
	from         string
	toList       []string
	ccList       []string
	subject      string
	content      string
	customArgs   map[string]string
	categories   []string
	inReplyTo    string
	referenceIds []string
}

func (m mailMsgImpl) From() string {
	return m.from
}

func (m mailMsgImpl) ToList() []string {
	return m.toList
}

func (m mailMsgImpl) CCList() []string {
	return m.ccList
}

func (m mailMsgImpl) Subject() string {
	return m.subject
}

func (m mailMsgImpl) Content() string {
	return m.content
}

func (m mailMsgImpl) CustomArgs() map[string]string {
	return m.customArgs
}

func (m mailMsgImpl) Categories() []string {
	return m.categories
}
func (m mailMsgImpl) InReplyTo() string {
	return m.inReplyTo
}
func (m mailMsgImpl) ReferenceIds() []string {
	return m.referenceIds
}

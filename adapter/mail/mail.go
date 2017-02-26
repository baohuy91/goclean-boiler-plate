package mail

// Mail struct that content message to feed to mail client to be sent away
type Mail interface {
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

type mailImpl struct {
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

func (m mailImpl) From() string {
	return m.from
}

func (m mailImpl) ToList() []string {
	return m.toList
}

func (m mailImpl) CCList() []string {
	return m.ccList
}

func (m mailImpl) Subject() string {
	return m.subject
}

func (m mailImpl) Content() string {
	return m.content
}

func (m mailImpl) CustomArgs() map[string]string {
	return m.customArgs
}

func (m mailImpl) Categories() []string {
	return m.categories
}
func (m mailImpl) InReplyTo() string {
	return m.inReplyTo
}
func (m mailImpl) ReferenceIds() []string {
	return m.referenceIds
}

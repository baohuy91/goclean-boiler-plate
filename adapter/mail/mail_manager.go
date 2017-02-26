package mail

// A mail client that sent the mail object away
type MailManager interface {
	SendMail(mail Mail) error
}

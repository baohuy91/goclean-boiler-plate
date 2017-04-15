package usecase

type MailService interface {
	SendMail(mailContent string, toUid string) error
}

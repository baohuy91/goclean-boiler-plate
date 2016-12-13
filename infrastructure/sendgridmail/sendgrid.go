package sendgridmail

import (
	"errors"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"strings"
	"github.com/sendgrid/sendgrid-go"
	"net/http"
	"goclean/interfaceadapter/controller"
)

type SendGridMailManager interface {
	SendMail(mail controller.Mail) error
}

func NewSendGridMailManager(host string, endPoint string, apiKey string) SendGridMailManager {
	return &MailManagerImpl{
		host:        host,
		endPoint:    endPoint,
		apiKey:      apiKey,
	}
}

type MailManagerImpl struct {
	host        string
	endPoint    string
	apiKey      string
}

// Send mail that can be label as reply
func (m *MailManagerImpl) SendMail(mailObj controller.Mail) error {
	if len(mailObj.ToList) == 0 {
		return errors.New("TO list can not be null")
	}
	toList := []*mail.Email{}
	for _, address := range mailObj.ToList() {
		toList = append(toList, mail.NewEmail("", address))
	}

	ccList := []*mail.Email{}
	for _, address := range mailObj.CCList() {
		ccList = append(ccList, mail.NewEmail("", address))
	}

	// SendGrid mail body
	sgMail := mail.NewV3Mail()
	sgMail.SetFrom(mail.NewEmail("", mailObj.From()))
	sgMail.AddContent(mail.NewContent("text/plain", mailObj.Content()))
	sgMail.Categories = append(sgMail.Categories, mailObj.Categories()...)
	for k, arg := range mailObj.CustomArgs() {
		sgMail.SetCustomArg(k, arg)
	}

	// personalize
	p := mail.NewPersonalization()
	p.AddTos(toList...)
	p.AddCCs(ccList...)
	p.SetHeader("In-Reply-To:", mailObj.InReplyTo())
	p.SetHeader("References:", strings.Join(mailObj.ReferenceIds(), " "))
	p.Subject = mailObj.Subject
	sgMail.AddPersonalizations(p)

	request := sendgrid.GetRequest(m.apiKey, m.endPoint, m.host)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(sgMail)

	resp, err := sendgrid.API(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return errors.New("Request Error Status code: "+ resp.StatusCode)
	}

	return nil
}

package notifier

import (
	"fmt"
	"net/smtp"
)

type gmailNotifier struct {
	gmail string
	auth  smtp.Auth
}

func NewGmailNotifier(gmail, password string) *gmailNotifier {
	return &gmailNotifier{
		auth:  smtp.PlainAuth("", gmail, password, "smtp.gmail.com"),
		gmail: gmail,
	}
}

func (n gmailNotifier) Notify(to string, title, body string) error {
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, title, body)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		n.auth,
		n.gmail,
		[]string{to},
		[]byte(message),
	)
}
package notifier

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type gmailNotifier struct {
	gmail          string
	gmailServer    string
	auth           smtp.Auth
	letterTemplate string
}

func NewGmailNotifier(gmailServer, gmail, password string) *gmailNotifier {
	letter, err := getGmailLetter("templates/index.html")
	if err != nil {
		panic(err)
	}

	return &gmailNotifier{
		auth:           smtp.PlainAuth("", gmail, password, gmailServer),
		gmail:          gmail,
		gmailServer:    gmailServer,
		letterTemplate: letter,
	}
}

func (n gmailNotifier) Notify(to string, title, body string) error {
	message := generateGmailLetter(n.letterTemplate, to, title, body)

	return smtp.SendMail(
		fmt.Sprintf("%s:587", n.gmailServer),
		n.auth,
		n.gmail,
		[]string{to},
		[]byte(message),
	)
}

func generateGmailLetter(template, to, title, body string) string {
	letter := strings.Replace(template, "{%BTCPRICE%}", body, -1)

	return "To: " + to + "\r\n" +
		"Subject: " + title + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		letter + "\r\n"
}

func getGmailLetter(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

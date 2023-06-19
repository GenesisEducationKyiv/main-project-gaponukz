package notifier

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type gmailNotifier struct {
	gmail       string
	gmailServer string
	auth        smtp.Auth
}

func NewGmailNotifier(gmailServer, gmail, password string) *gmailNotifier {
	return &gmailNotifier{
		auth:        smtp.PlainAuth("", gmail, password, gmailServer),
		gmail:       gmail,
		gmailServer: gmailServer,
	}
}

func (n gmailNotifier) Notify(to string, title, body string) error {
	letter, err := getGmailLetter("templates/index.html", body)
	if err != nil {
		return err
	}

	message := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + title + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			letter + "\r\n",
	)
	return smtp.SendMail(
		fmt.Sprintf("%s:587", n.gmailServer),
		n.auth,
		n.gmail,
		[]string{to},
		[]byte(message),
	)
}

func getGmailLetter(filename string, content string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return strings.Replace(string(data), "{%BTCPRICE%}", content, -1), nil
}

package observer

import (
	"btcapp/src/settings"
	"fmt"
	"net/smtp"
)

func SendGmailLetterStrategy(
	userGmail string,
	BTCPrice float64,
	settings settings.Settings,
) error {
	message := fmt.Sprintf("To: %s\r\nSubject: BTC/UAH price\r\n\r\nBTC/UAH price: %f\r\n", userGmail, BTCPrice)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth(
			"",
			settings.Gmail,
			settings.GmailPassword,
			"smtp.gmail.com",
		),
		settings.Gmail,
		[]string{userGmail},
		[]byte(message),
	)
}

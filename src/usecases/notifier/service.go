package notifier

import (
	"btcapp/src/entities"
	"fmt"
)

type notifier interface {
	Notify(to string, title, body string) error
}

type notifierService struct {
	n notifier
}

func (s *notifierService) NotifyBTCPrice(users []entities.User, price entities.Price) {
	var (
		title = "BTC price update!"
		body  = fmt.Sprintf("%f", price)
	)

	for _, user := range users {
		s.n.Notify(user.Gmail, title, body)
	}
}

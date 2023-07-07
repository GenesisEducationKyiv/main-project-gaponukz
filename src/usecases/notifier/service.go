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

func NewNotifierService(n notifier) notifierService {
	return notifierService{n: n}
}

func (s notifierService) NotifyBTCPrice(users []entities.User, price entities.Price) {
	var (
		title = "BTC price update!"
		body  = fmt.Sprintf("%f", price)
	)

	for _, user := range users {
		err := s.n.Notify(user.Gmail, title, body)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

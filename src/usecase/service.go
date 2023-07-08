package usecase

import (
	"btcapp/src/entities"
	"fmt"
)

type userStorage interface {
	GetAll() ([]entities.User, error)
	Create(entities.User) error
}

type rateExporter interface {
	CurrentBTCPrice() (float64, error)
}

type notier interface {
	Notify(to string, title, body string) error
}

type service struct {
	storage  userStorage
	exporter rateExporter
	notier   notier
}

func NewService(s userStorage, e rateExporter, n notier) *service {
	return &service{s, e, n}
}

func (s service) GetCurrentPrice() (float64, error) {
	return s.exporter.CurrentBTCPrice()
}

func (s service) SubscribeUser(user entities.User) error {
	return s.storage.Create(user)
}

func (s service) NotifySubscribers() error {
	users, err := s.storage.GetAll()
	if err != nil {
		return err
	}

	btcPrice, err := s.exporter.CurrentBTCPrice()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := s.notier.Notify(user.Gmail, "BTC/UAH price", fmt.Sprintf("%f", btcPrice))
		if err != nil {
			fmt.Printf("Could not send to %s: %s", user.Gmail, err.Error())
		}
	}

	return nil
}

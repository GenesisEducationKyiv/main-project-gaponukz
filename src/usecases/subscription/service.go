package subscription

import (
	"btcapp/src/entities"
	"errors"
)

type userStorage interface {
	GetAll() ([]entities.User, error)
	Create(entities.User) error
}

type subscriptionService struct {
	us userStorage
}

func NewSubscriptionService(us userStorage) subscriptionService {
	return subscriptionService{us: us}
}

func (s subscriptionService) All() ([]entities.User, error) {
	return s.us.GetAll()
}

func (s subscriptionService) SubscribeUser(user entities.User) error {
	if s.isUserAlreadySubscribed(user) {
		return errors.New("user already subscribed")
	}

	return s.us.Create(user)
}

func (s subscriptionService) isUserAlreadySubscribed(user entities.User) bool {
	users, _ := s.us.GetAll()

	for _, u := range users {
		if u.Gmail == user.Gmail {
			return true
		}
	}

	return false
}

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

func (s subscriptionService) All(filter func(entities.User) bool) ([]entities.User, error) {
	if filter == nil {
		filter = func(user entities.User) bool { return true }
	}

	users, err := s.us.GetAll()
	if err != nil {
		return nil, err
	}

	var filtered []entities.User

	for _, user := range users {
		if filter(user) {
			filtered = append(filtered, user)
		}
	}

	return filtered, nil
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

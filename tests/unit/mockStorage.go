package unit

import "btcapp/src/entities"

type storageMock struct {
	users []entities.User
}

func (sm *storageMock) GetAll() ([]entities.User, error) {
	return sm.users, nil
}

func (sm *storageMock) Create(user entities.User) error {
	sm.users = append(sm.users, user)
	return nil
}

func (sm *storageMock) Contains(user entities.User) bool {
	for _, other := range sm.users {
		if other.Gmail == other.Gmail {
			return true
		}
	}

	return false
}

func NewStorageMock() *storageMock {
	return &storageMock{}
}

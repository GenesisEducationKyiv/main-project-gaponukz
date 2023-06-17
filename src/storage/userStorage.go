package storage

import "btcapp/src/entities"

type IUserStorage interface {
	GetAll() ([]entities.User, error)
	Create(entities.User) error
}

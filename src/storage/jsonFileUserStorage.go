package storage

import (
	"btcapp/src/entities"
	"encoding/json"
	"io"
	"os"
)

type jsonFileUserStorage struct {
	Filename string
}

func NewJsonFileUserStorage(filename string) *jsonFileUserStorage {
	return &jsonFileUserStorage{Filename: filename}
}

func (s jsonFileUserStorage) GetAll() ([]entities.User, error) {
	jsonFile, err := os.Open(s.Filename)

	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	users := []entities.User{}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(byteValue, &users)

	return users, nil
}

func (s jsonFileUserStorage) Create(user entities.User) error {
	users, err := s.GetAll()

	if err != nil {
		return err
	}

	users = append(users, user)
	err = s.writeUsers(users)

	return err
}

func (s jsonFileUserStorage) writeUsers(users []entities.User) error {
	usersJSON, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	file, err := os.Create(s.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(usersJSON)
	if err != nil {
		return err
	}

	return nil
}

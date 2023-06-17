package storage

import (
	"btcapp/src/entities"
	"encoding/json"
	"io"
	"os"
)

type JsonFileUserStorage struct {
	Filename string
}

func NewJsonFileUserStorage(filename string) *JsonFileUserStorage {
	return &JsonFileUserStorage{Filename: filename}
}

func (s JsonFileUserStorage) GetAll() ([]entities.User, error) {
	jsonFile, err := os.Open(s.Filename)

	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	users := []entities.User{}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)

	return users, nil
}

func (s JsonFileUserStorage) Create(user entities.User) error {
	users, err := s.GetAll()

	if err != nil {
		return err
	}

	users = append(users, user)
	err = s.writeUsers(users)

	return err
}

func (s JsonFileUserStorage) writeUsers(users []entities.User) error {
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

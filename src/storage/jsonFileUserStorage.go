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

func (strg JsonFileUserStorage) GetAll() ([]entities.User, error) {
	jsonFile, err := os.Open(strg.Filename)

	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	users := []entities.User{}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)

	return users, nil
}

func (strg JsonFileUserStorage) Create(user entities.User) error {
	users, err := strg.GetAll()

	if err != nil {
		return err
	}

	users = append(users, user)
	err = strg.writeUsers(users)

	return err
}

func (strg JsonFileUserStorage) writeUsers(users []entities.User) error {
	usersJSON, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	file, err := os.Create(strg.Filename)
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

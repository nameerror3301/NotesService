package models

import (
	crypt "NotesService/tools"

	"github.com/sirupsen/logrus"
)

type UserData struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

// Хранилеще пользователей
var userData = map[string][]byte{}

// Проверка аутентификации
func IsUserData(email string, pass string) bool {
	hashPass, isOk := userData[email]
	if !isOk {
		return false
	}

	if checkHash := crypt.CheckControlSum(pass, string(hashPass)); !checkHash {
		return false
	}
	return true
}

// Создание пользователя
func CreateUser(email string, pass string) (bool, error) {
	for key := range userData {
		if key == email {
			return false, nil
		}
	}
	if hashPass, err := crypt.HashingPassword(pass); err != nil {
		logrus.Debugf("Err hashing password - %s", err)
		return false, err
	} else {
		userData[email] = hashPass
	}

	return true, nil
}

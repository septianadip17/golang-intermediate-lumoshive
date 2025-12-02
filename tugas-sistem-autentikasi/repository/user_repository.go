package repository

import (
	"encoding/json"
	"os"
	"strings"

	"sistem-autentikasi/models"
)

const dbFile = "storage/users.json"

func LoadUsers() ([]models.User, error) {
	f, err := os.Open(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.User{}, nil
		}
		return nil, err
	}
	defer f.Close()

	var users []models.User
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&users); err != nil {
		return []models.User{}, nil
	}

	return users, nil
}

func SaveUsers(users []models.User) error {
	f, err := os.Create(dbFile)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	return encoder.Encode(users)
}

func FindUserByEmail(users []models.User, email string) (models.User, bool) {
	for _, u := range users {
		if strings.EqualFold(u.Email, email) {
			return u, true
		}
	}
	return models.User{}, false
}

func EmailExists(users []models.User, email string) bool {
	_, found := FindUserByEmail(users, email)
	return found
}

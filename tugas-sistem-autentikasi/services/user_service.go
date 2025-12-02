package services

import (
	"errors"

	"sistem-autentikasi/dto"
	"sistem-autentikasi/models"
	"sistem-autentikasi/repository"
	"sistem-autentikasi/utils"
)

type UserService struct{}

func (UserService) Register(input dto.RegisterDTO) error {
	if !utils.ValidEmail(input.Email) {
		return errors.New("email tidak valid")
	}

	if !utils.ValidPhone(input.PhoneNumber) {
		return errors.New("nomor hp tidak valid")
	}

	if !utils.ValidPassword(input.Password) {
		return errors.New("password minimal 6 karakter")
	}

	users, err := repository.LoadUsers()
	if err != nil {
		return err
	}

	if repository.EmailExists(users, input.Email) {
		return errors.New("email sudah terdaftar")
	}

	newUser := models.User{
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	users = append(users, newUser)
	return repository.SaveUsers(users)
}

func (UserService) Login(input dto.LoginDTO) (models.User, error) {
	users, err := repository.LoadUsers()
	if err != nil {
		return models.User{}, err
	}

	user, found := repository.FindUserByEmail(users, input.Email)
	if !found {
		return models.User{}, errors.New("email tidak ditemukan")
	}

	if user.Password != input.Password {
		return models.User{}, errors.New("password salah")
	}

	return user, nil
}

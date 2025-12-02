package handlers

import (
	"bufio"
	"fmt"

	"sistem-autentikasi/dto"
	"sistem-autentikasi/services"
	"sistem-autentikasi/utils"
)

type UserHandler struct {
	Service services.UserService
}

func (h UserHandler) Register(scanner *bufio.Scanner) {
	email, _ := utils.ReadLine(scanner, "Email: ")
	phone, _ := utils.ReadLine(scanner, "Phone number: ")
	password, _ := utils.ReadLine(scanner, "Password: ")

	input := dto.RegisterDTO{
		Email:       email,
		PhoneNumber: phone,
		Password:    password,
	}

	err := h.Service.Register(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Registrasi berhasil")
}

func (h UserHandler) Login(scanner *bufio.Scanner) {
	email, _ := utils.ReadLine(scanner, "Email: ")
	password, _ := utils.ReadLine(scanner, "Password: ")

	input := dto.LoginDTO{
		Email:    email,
		Password: password,
	}

	user, err := h.Service.Login(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Login berhasil. Selamat datang", user.Email)
}

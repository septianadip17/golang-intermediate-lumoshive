package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type User struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

const dbFile = "users.json"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("1) Register")
		fmt.Println("2) Login")
		fmt.Println("3) Exit")
		fmt.Print("Pilih: ")

		if !scanner.Scan() {
			return
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			if err := handleRegister(scanner); err != nil {
				fmt.Println("Error:", err)
			}
		case "2":
			if err := handleLogin(scanner); err != nil {
				fmt.Println("Error:", err)
			}
		case "3":
			fmt.Println("Selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func handleRegister(scanner *bufio.Scanner) error {
	fmt.Print("Email: ")
	if !scanner.Scan() {
		return errors.New("input gagal")
	}
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Phone number: ")
	if !scanner.Scan() {
		return errors.New("input gagal")
	}
	phone := strings.TrimSpace(scanner.Text())

	fmt.Print("Password: ")
	if !scanner.Scan() {
		return errors.New("input gagal")
	}
	password := strings.TrimSpace(scanner.Text())

	if err := validateRegisterInput(email, phone, password); err != nil {
		return err
	}

	users, err := loadUsers()
	if err != nil {
		return err
	}

	if emailExists(users, email) {
		return errors.New("email sudah terdaftar")
	}

	newUser := User{
		Email:       email,
		PhoneNumber: phone,
		Password:    password, // tidak hash
	}

	users = append(users, newUser)

	if err := saveUsers(users); err != nil {
		return err
	}

	fmt.Println("Registrasi berhasil")
	return nil
}

func handleLogin(scanner *bufio.Scanner) error {
	fmt.Print("Email: ")
	if !scanner.Scan() {
		return errors.New("input gagal")
	}
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Password: ")
	if !scanner.Scan() {
		return errors.New("input gagal")
	}
	password := strings.TrimSpace(scanner.Text())

	users, err := loadUsers()
	if err != nil {
		return err
	}

	user, found := findUserByEmail(users, email)
	if !found {
		return errors.New("email tidak ditemukan")
	}

	if user.Password != password {
		return errors.New("password salah")
	}

	fmt.Println("Login berhasil. Selamat datang", user.Email)
	return nil
}

func validateRegisterInput(email, phone, password string) error {
	if !validEmail(email) {
		return errors.New("email tidak valid")
	}
	if !validPhone(phone) {
		return errors.New("phone number tidak valid. Hanya angka. Panjang 10-15 digit")
	}
	if len(password) < 6 {
		return errors.New("password minimal 6 karakter")
	}
	return nil
}

func validEmail(email string) bool {
	re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	return re.MatchString(email)
}

func validPhone(phone string) bool {
	re := regexp.MustCompile(`^[0-9]{10,15}$`)
	return re.MatchString(phone)
}

func loadUsers() ([]User, error) {
	f, err := os.Open(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	}
	defer f.Close()

	var users []User
	dec := json.NewDecoder(f)

	if err := dec.Decode(&users); err != nil {
		return []User{}, nil
	}

	return users, nil
}

func saveUsers(users []User) error {
	f, err := os.Create(dbFile)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(users)
}

func emailExists(users []User, email string) bool {
	for _, u := range users {
		if strings.EqualFold(u.Email, email) {
			return true
		}
	}
	return false
}

func findUserByEmail(users []User, email string) (User, bool) {
	for _, u := range users {
		if strings.EqualFold(u.Email, email) {
			return u, true
		}
	}
	return User{}, false
}

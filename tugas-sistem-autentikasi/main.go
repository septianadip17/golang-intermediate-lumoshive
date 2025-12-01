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

// Baca input satu baris dengan prompt
func readLine(scanner *bufio.Scanner, prompt string) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return "", errors.New("input gagal")
	}
	return strings.TrimSpace(scanner.Text()), nil
}

// Minta email sampai valid format
func promptEmail(scanner *bufio.Scanner) (string, error) {
	for {
		email, err := readLine(scanner, "Email: ")
		if err != nil {
			return "", err
		}
		if !validEmail(email) {
			fmt.Println("Email tidak valid. Format harus ada @ dan domain.")
			continue
		}
		return email, nil
	}
}

// Minta phone sampai valid
func promptPhone(scanner *bufio.Scanner) (string, error) {
	for {
		phone, err := readLine(scanner, "Phone number: ")
		if err != nil {
			return "", err
		}
		if !validPhone(phone) {
			fmt.Println("Nomor tidak valid. Hanya angka. Panjang 10-15.")
			continue
		}
		return phone, nil
	}
}

// Minta password sampai valid
func promptPassword(scanner *bufio.Scanner) (string, error) {
	for {
		pass, err := readLine(scanner, "Password: ")
		if err != nil {
			return "", err
		}
		if len(pass) < 6 {
			fmt.Println("Password minimal 6 karakter.")
			continue
		}
		return pass, nil
	}
}

func handleRegister(scanner *bufio.Scanner) error {
	// Input dan validasi bertahap. Tidak kembali ke menu sampai valid.
	email, err := promptEmail(scanner)
	if err != nil {
		return err
	}

	phone, err := promptPhone(scanner)
	if err != nil {
		return err
	}

	password, err := promptPassword(scanner)
	if err != nil {
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
		Password:    password, // disimpan plain text sesuai permintaan
	}

	users = append(users, newUser)

	if err := saveUsers(users); err != nil {
		return err
	}

	fmt.Println("Registrasi berhasil")
	return nil
}

func handleLogin(scanner *bufio.Scanner) error {
	// Terus minta email + password sampai login berhasil.
	for {
		email, err := readLine(scanner, "Email: ")
		if err != nil {
			return err
		}
		password, err := readLine(scanner, "Password: ")
		if err != nil {
			return err
		}

		users, err := loadUsers()
		if err != nil {
			return err
		}

		user, found := findUserByEmail(users, email)
		if !found {
			fmt.Println("Email tidak ditemukan. Coba lagi.")
			continue
		}

		if user.Password != password {
			fmt.Println("Password salah. Coba lagi.")
			continue
		}

		fmt.Println("Login berhasil. Selamat datang", user.Email)
		return nil
	}
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
		// Jika file kosong atau rusak, kembalikan slice kosong
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

package main

import (
	"bufio"
	"fmt"
	"os"

	"sistem-autentikasi/handlers"
	"sistem-autentikasi/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	userService := services.UserService{}
	userHandler := handlers.UserHandler{Service: userService}

	for {
		fmt.Println()
		fmt.Println("1) Register")
		fmt.Println("2) Login")
		fmt.Println("3) Exit")
		fmt.Print("Pilih: ")

		if !scanner.Scan() {
			return
		}
		choice := scanner.Text()

		switch choice {
		case "1":
			userHandler.Register(scanner)
		case "2":
			userHandler.Login(scanner)
		case "3":
			fmt.Println("Selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

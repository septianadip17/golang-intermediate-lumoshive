package utils

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func ReadLine(scanner *bufio.Scanner, prompt string) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return "", errors.New("gagal membaca input")
	}
	return strings.TrimSpace(scanner.Text()), nil
}

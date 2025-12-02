package utils

import "regexp"

func ValidEmail(email string) bool {
	re := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	return re.MatchString(email)
}

func ValidPhone(phone string) bool {
	re := regexp.MustCompile(`^[0-9]{10,15}$`)
	return re.MatchString(phone)
}

func ValidPassword(password string) bool {
	return len(password) >= 6
}

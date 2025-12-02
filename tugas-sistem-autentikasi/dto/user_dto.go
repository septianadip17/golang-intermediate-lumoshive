package dto

type RegisterDTO struct {
	Email       string
	PhoneNumber string
	Password    string
}

type LoginDTO struct {
	Email    string
	Password string
}

package dtos

import "go_pro/internal/validations"

type UserRequest struct {
	Email    string
	Password string
	validations.FormValidator
}

func NewUserRequest(email, password string) (req UserRequest) {
	req.Email = email
	req.Password = password

	return
}

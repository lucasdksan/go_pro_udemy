package handlers

import (
	"go_pro/internal/dtos"
	"go_pro/tools"
	"net/http"
	"strings"
)

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (uh *userHandler) SignupForm(w http.ResponseWriter, r *http.Request) error {
	return render(w, "user-signup.html", nil, http.StatusOK)
}

func (uh *userHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := dtos.NewUserRequest(email, password)

	if strings.TrimSpace(user.Password) == "" {
		user.AddFieldError("password", "Senha é obrigatória")
	}

	if !user.Valid() {
		render(w, "user-signup.html", user, http.StatusUnprocessableEntity)
		return nil
	}

	if err := tools.ValidateEmail(user.Email); err != nil {
		user.AddFieldError("email", "Email é inválido")
	}

	return render(w, "user-signup.html", nil, http.StatusOK)
}

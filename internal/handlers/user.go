package handlers

import (
	"go_pro/internal/dtos"
	"go_pro/internal/repositories"
	"go_pro/tools"
	"net/http"
	"strings"
)

type userHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *userHandler {
	return &userHandler{repo: repo}
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

	user_response, err := uh.repo.Create(r.Context(), user.Email, user.Password)

	if err == repositories.ErrDuplicateEmail {
		user.AddFieldError("email", "email já está cadastrado")
		return render(w, "user-signup.html", user, http.StatusUnprocessableEntity)
	}

	if err != nil {
		return err
	}

	return render(w, "user-signup-success.html", user_response, http.StatusOK)
}

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

	hash, err := tools.HashPassword(user.Password)

	if err != nil {
		return err
	}

	hashKey := tools.GenerateToken()

	_, token, err := uh.repo.Create(r.Context(), user.Email, hash, hashKey)

	if err == repositories.ErrDuplicateEmail {
		user.AddFieldError("email", "email já está cadastrado")
		return render(w, "user-signup.html", user, http.StatusUnprocessableEntity)
	}

	if err != nil {
		return err
	}

	return render(w, "user-signup-success.html", token, http.StatusOK)
}

func (uh *userHandler) Confirm(w http.ResponseWriter, r *http.Request) error {
	token := r.PathValue("token")
	msg := "Seu cadastro foi confirmado. Agora você já pode fazer o login no sistema."

	if err := uh.repo.ConfirmUserByToken(r.Context(), token); err != nil {
		msg = "Esse cadastro já foi confirmado ou o token é inválid"
	}

	return render(w, "user-confirm.html", msg, http.StatusOK)
}

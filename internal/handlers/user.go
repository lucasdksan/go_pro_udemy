package handlers

import (
	"fmt"
	"go_pro/internal/dtos"
	"go_pro/internal/repositories"
	"go_pro/tools"
	"net/http"
	"strings"

	"github.com/alexedwards/scs/v2"
)

type userHandler struct {
	session *scs.SessionManager
	repo    repositories.UserRepository
}

func NewUserHandler(session *scs.SessionManager, repo repositories.UserRepository) *userHandler {
	return &userHandler{repo: repo, session: session}
}

func (uh *userHandler) SigninForm(w http.ResponseWriter, r *http.Request) error {
	userId := uh.session.GetInt64(r.Context(), "userId")

	fmt.Println("USER ID: ", userId)

	return render(w, r, "user-signin.html", nil, http.StatusOK)
}

func (uh *userHandler) Signin(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := dtos.NewUserRequest(email, password)

	if strings.TrimSpace(user.Password) == "" {
		user.AddFieldError("password", "Senha é obrigatória")
	}

	if err := tools.ValidateEmail(user.Email); err != nil {
		user.AddFieldError("email", "Email é inválido")
	}

	if !user.Valid() {
		render(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
		return nil
	}

	data, err := uh.repo.FindByEmail(r.Context(), user.Email)

	if err != nil {
		user.AddFieldError("validation", "invalid credentials")
		return render(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
	}

	if !data.Active.Bool {
		user.AddFieldError("validation", "user did not confirm registration")
		return render(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
	}

	if !tools.ValidatePassword(user.Password, data.Password.String) {
		user.AddFieldError("validation", "invalid credentials")
		return render(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
	}

	uh.session.Put(r.Context(), "userId", data.Id.Int.Int64())

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
	return nil
}

func (uh *userHandler) Me(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("session")

	if err != nil {
		http.Redirect(w, r, "/user/signin", http.StatusTemporaryRedirect)
		return nil
	}

	fmt.Fprintf(w, "Email: %s", cookie.Value)
	return nil
}

func (uh *userHandler) SignupForm(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, "user-signup.html", nil, http.StatusOK)
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

	if len(strings.TrimSpace(user.Password)) < 6 {
		user.AddFieldError("password", "Senha precisa ter no mínimo 6 caracteres")
	}

	if err := tools.ValidateEmail(user.Email); err != nil {
		user.AddFieldError("email", "Email é inválido")
	}

	if !user.Valid() {
		render(w, r, "user-signup.html", user, http.StatusUnprocessableEntity)
		return nil
	}

	hash, err := tools.HashPassword(user.Password)

	if err != nil {
		return err
	}

	hashKey := tools.GenerateToken()

	_, token, err := uh.repo.Create(r.Context(), user.Email, hash, hashKey)

	if err == repositories.ErrDuplicateEmail {
		user.AddFieldError("email", "email já está cadastrado")
		return render(w, r, "user-signup.html", user, http.StatusUnprocessableEntity)
	}

	if err != nil {
		return err
	}

	return render(w, r, "user-signup-success.html", token, http.StatusOK)
}

func (uh *userHandler) Confirm(w http.ResponseWriter, r *http.Request) error {
	token := r.PathValue("token")
	msg := "Seu cadastro foi confirmado. Agora você já pode fazer o login no sistema."

	if err := uh.repo.ConfirmUserByToken(r.Context(), token); err != nil {
		msg = "Esse cadastro já foi confirmado ou o token é inválid"
	}

	return render(w, r, "user-confirm.html", msg, http.StatusOK)
}

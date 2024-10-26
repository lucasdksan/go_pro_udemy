package handlers

import (
	"fmt"
	"go_pro/internal/dtos"
	"go_pro/internal/mailers"
	"go_pro/internal/render"
	"go_pro/internal/repositories"
	"go_pro/tools"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type userHandler struct {
	render  *render.RenderTemplate
	session *scs.SessionManager
	repo    repositories.UserRepository
	mail    mailers.MailService
}

func NewUserHandler(render *render.RenderTemplate, session *scs.SessionManager, repo repositories.UserRepository, mail mailers.MailService) *userHandler {
	return &userHandler{repo: repo, session: session, render: render, mail: mail}
}

func (uh *userHandler) SigninForm(w http.ResponseWriter, r *http.Request) error {
	userId := uh.session.GetInt64(r.Context(), "userId")

	fmt.Println("USER ID: ", userId)

	return uh.render.RenderPage(w, r, "user-signin.html", nil, http.StatusOK)
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
		uh.render.RenderPage(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
		return nil
	}

	data, err := uh.repo.FindByEmail(r.Context(), user.Email)

	if err != nil {
		user.AddFieldError("validation", "invalid credentials")
		return uh.render.RenderPage(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
	}

	if !data.Active.Bool {
		user.AddFieldError("validation", "user did not confirm registration")
		return uh.render.RenderPage(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
	}

	if !tools.ValidatePassword(user.Password, data.Password.String) {
		user.AddFieldError("validation", "invalid credentials")
		return uh.render.RenderPage(w, r, "user-signin.html", user, http.StatusUnprocessableEntity)
	}

	if err := uh.session.RenewToken(r.Context()); err != nil {
		slog.Error(err.Error())
		return err
	}

	uh.session.Put(r.Context(), "userId", data.Id.Int.Int64())
	uh.session.Put(r.Context(), "userEmail", data.Email.String)

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
	data := dtos.UserRequest{}
	data.Flash = uh.session.PopString(r.Context(), "flash")

	return uh.render.RenderPage(w, r, "user-signup.html", data, http.StatusOK)
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
		uh.render.RenderPage(w, r, "user-signup.html", user, http.StatusUnprocessableEntity)
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
		return uh.render.RenderPage(w, r, "user-signup.html", user, http.StatusUnprocessableEntity)
	}

	if err != nil {
		return err
	}

	body, err := uh.render.RenderMailBody(r, "confirmation.html", map[string]string{"token": token})

	if err != nil {
		return err
	}

	if err := uh.mail.Send(mailers.MailMessage{
		To:      []string{user.Email},
		Subject: "Confirmação de Cadastro",
		IsHTML:  true,
		Body:    body,
	}); err != nil {
		slog.Error(err.Error())
		return err
	}

	return uh.render.RenderPage(w, r, "user-signup-success.html", token, http.StatusOK)
}

func (uh *userHandler) Confirm(w http.ResponseWriter, r *http.Request) error {
	token := r.PathValue("token")
	msg := "Seu cadastro foi confirmado. Agora você já pode fazer o login no sistema."

	if err := uh.repo.ConfirmUserByToken(r.Context(), token); err != nil {
		msg = "Esse cadastro já foi confirmado ou o token é inválid"
	}

	return uh.render.RenderPage(w, r, "user-confirm.html", msg, http.StatusOK)
}

func (uh *userHandler) ForgetPasswordForm(w http.ResponseWriter, r *http.Request) error {
	return uh.render.RenderPage(w, r, "user-forget-password.html", nil, http.StatusOK)
}

func (uh *userHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) error {
	email := r.PostFormValue("email")
	hashToken := tools.GenerateToken()

	token, err := uh.repo.CreateResetPasswordToken(r.Context(), email, hashToken)

	if err != nil {
		data := dtos.UserRequest{}
		data.Email = email
		data.AddFieldError("email", "Email não possui cadastro válido no sistema")
		return uh.render.RenderPage(w, r, "user-forget-password.html", data, http.StatusOK)
	}

	body, err := uh.render.RenderMailBody(r, "forgetpassword.html", map[string]string{"token": token})

	if err != nil {
		return err
	}

	err = uh.mail.Send(mailers.MailMessage{
		To:      []string{email},
		Subject: "Resetar senha",
		IsHTML:  true,
		Body:    body,
	})

	if err != nil {
		return err
	}

	message := "Foi enviado um email com um link para que você possa resetar a sua senha."

	return uh.render.RenderPage(w, r, "generic-success.html", message, http.StatusOK)
}

func (uh *userHandler) Signout(w http.ResponseWriter, r *http.Request) error {
	if err := uh.session.RenewToken(r.Context()); err != nil {
		slog.Error(err.Error())
		return err
	}

	uh.session.Remove(r.Context(), "userId")

	http.Redirect(w, r, "/user/signin", http.StatusSeeOther)
	return nil
}

func (uh *userHandler) ResetPasswordForm(w http.ResponseWriter, r *http.Request) error {
	token := r.PathValue("token")

	userToken, err := uh.repo.GetUserConfirmationByToken(r.Context(), token)
	elapsedTime := time.Since(userToken.CreatedAt.Time).Hours()

	if err != nil || userToken.Confirmed.Bool || elapsedTime > 4 {
		msg := "Token inválido ou expirado"
		return uh.render.RenderPage(w, r, "generic-error.html", msg, http.StatusOK)
	}

	data := struct {
		Token  string
		Errors []string
	}{
		Token: token,
	}

	return uh.render.RenderPage(w, r, "user-reset-password.html", data, http.StatusOK)
}

func (uh *userHandler) ResetPassword(w http.ResponseWriter, r *http.Request) error {
	pass := r.PostFormValue("password")
	token := r.PostFormValue("token")

	hashedPass, err := tools.HashPassword(pass)

	if err != nil {
		data := struct {
			Token  string
			Errors []string
		}{
			Token:  token,
			Errors: []string{"não foi possível alterar a senha"},
		}
		return uh.render.RenderPage(w, r, "user-reset-password.html", data, http.StatusOK)
	}

	email, err := uh.repo.UpdatePasswordByToken(r.Context(), hashedPass, token)

	if err != nil {
		data := struct {
			Token  string
			Errors []string
		}{
			Token:  token,
			Errors: []string{"não foi possível alterar a senha. Solicite uma nova"},
		}

		return uh.render.RenderPage(w, r, "user-reset-password.html", data, http.StatusOK)
	}

	uh.session.Put(r.Context(), "flash", "Sua senha foi atualizada. Agora você pode fazer o login")

	uh.mail.Send(mailers.MailMessage{
		To:      []string{email},
		Subject: "Sua senha foi atualizada",
		Body:    []byte("Sua senha foi atualizada e agora você já pode fazer o login novamente"),
		IsHTML:  false,
	})

	http.Redirect(w, r, "user/signin", http.StatusSeeOther)
	return nil
}

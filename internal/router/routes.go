package router

import (
	"go_pro/assets"
	"go_pro/internal/handlers"
	"go_pro/internal/mailers"
	"go_pro/internal/render"
	"go_pro/internal/repositories"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadRoutes(sessionManager *scs.SessionManager, db *pgxpool.Pool, noteRepo repositories.NoteRepository, userRepo repositories.UserRepository, mail mailers.MailService) http.Handler {
	static, err := fs.Sub(assets.Files, ".")

	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	mux := http.NewServeMux()
	render := render.NewRender(sessionManager)
	staticHandler := http.FileServerFS(static)
	noteHandlers := handlers.NewNoteHandler(render, sessionManager, noteRepo)
	userHandlers := handlers.NewUserHandler(render, sessionManager, userRepo, mail)
	authMidd := handlers.NewAuthMiddleware(sessionManager)
	errorMidd := handlers.NewErrorHandlerMiddleware(render)

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", staticHandler))

	mux.HandleFunc("GET /", handlers.NewHomeHandler(render).HomeHandler)

	mux.Handle("GET /notes", authMidd.RequireAuth(errorMidd.HandlerError(noteHandlers.NoteList)))
	mux.Handle("GET /notes/{id}", authMidd.RequireAuth(errorMidd.HandlerError(noteHandlers.NoteView)))
	mux.Handle("GET /notes/new", authMidd.RequireAuth(errorMidd.HandlerError(noteHandlers.NoteNew)))
	mux.Handle("POST /notes", authMidd.RequireAuth(errorMidd.HandlerError(noteHandlers.NoteSave)))
	mux.Handle("DELETE /notes/{id}", authMidd.RequireAuth(errorMidd.HandlerError(noteHandlers.NoteDelete)))
	mux.Handle("GET /notes/{id}/update", authMidd.RequireAuth(errorMidd.HandlerError(noteHandlers.NoteEdit)))

	mux.Handle("GET /user/signup", errorMidd.HandlerError(userHandlers.SignupForm))
	mux.Handle("POST /user/signup", errorMidd.HandlerError(userHandlers.Signup))
	mux.Handle("GET /user/signin", errorMidd.HandlerError(userHandlers.SigninForm))
	mux.Handle("POST /user/signin", errorMidd.HandlerError(userHandlers.Signin))
	mux.Handle("GET /user/signout", errorMidd.HandlerError(userHandlers.Signout))
	mux.Handle("GET /user/forgetpassword", errorMidd.HandlerError(userHandlers.ForgetPasswordForm))
	mux.Handle("POST /user/forgetpassword", errorMidd.HandlerError(userHandlers.ForgetPassword))
	mux.Handle("POST /user/password", errorMidd.HandlerError(userHandlers.ResetPassword))
	mux.Handle("GET /user/password/{token}", errorMidd.HandlerError(userHandlers.ResetPasswordForm))

	mux.Handle("GET /confirmation/{token}", errorMidd.HandlerError(userHandlers.Confirm))

	mux.Handle("GET /me", errorMidd.HandlerError(userHandlers.Me))

	return mux
}

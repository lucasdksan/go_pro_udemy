package router

import (
	"go_pro/internal/handlers"
	"go_pro/internal/mailers"
	"go_pro/internal/render"
	"go_pro/internal/repositories"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadRoutes(sessionManager *scs.SessionManager, db *pgxpool.Pool, noteRepo repositories.NoteRepository, userRepo repositories.UserRepository, mail mailers.MailService) http.Handler {
	mux := http.NewServeMux()
	render := render.NewRender(sessionManager)
	staticHandler := http.FileServer(http.Dir("./assets/"))
	noteHandlers := handlers.NewNoteHandler(render, sessionManager, noteRepo)
	userHandlers := handlers.NewUserHandler(render, sessionManager, userRepo, mail)
	authMidd := handlers.NewAuthMiddleware(sessionManager)

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", staticHandler))

	mux.HandleFunc("GET /", handlers.NewHomeHandler(render).HomeHandler)

	mux.Handle("GET /notes", authMidd.RequireAuth(handlers.HandlerWithError(noteHandlers.NoteList)))
	mux.Handle("GET /notes/{id}", authMidd.RequireAuth(handlers.HandlerWithError(noteHandlers.NoteView)))
	mux.Handle("GET /notes/new", authMidd.RequireAuth(handlers.HandlerWithError(noteHandlers.NoteNew)))
	mux.Handle("POST /notes", authMidd.RequireAuth(handlers.HandlerWithError(noteHandlers.NoteSave)))
	mux.Handle("DELETE /notes/{id}", authMidd.RequireAuth(handlers.HandlerWithError(noteHandlers.NoteDelete)))
	mux.Handle("GET /notes/{id}/update", authMidd.RequireAuth(handlers.HandlerWithError(noteHandlers.NoteEdit)))

	mux.Handle("GET /user/signup", handlers.HandlerWithError(userHandlers.SignupForm))
	mux.Handle("POST /user/signup", handlers.HandlerWithError(userHandlers.Signup))
	mux.Handle("GET /user/signin", handlers.HandlerWithError(userHandlers.SigninForm))
	mux.Handle("POST /user/signin", handlers.HandlerWithError(userHandlers.Signin))
	mux.Handle("GET /user/signout", handlers.HandlerWithError(userHandlers.Signout))
	mux.Handle("GET /user/forgetpassword", handlers.HandlerWithError(userHandlers.ForgetPasswordForm))
	mux.Handle("POST /user/forgetpassword", handlers.HandlerWithError(userHandlers.ForgetPassword))
	mux.Handle("POST /user/password", handlers.HandlerWithError(userHandlers.ResetPassword))
	mux.Handle("GET /user/password/{token}", handlers.HandlerWithError(userHandlers.ResetPasswordForm))

	mux.Handle("GET /confirmation/{token}", handlers.HandlerWithError(userHandlers.Confirm))

	mux.Handle("GET /me", handlers.HandlerWithError(userHandlers.Me))

	return mux
}

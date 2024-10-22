package main

import (
	"fmt"
	"go_pro/config"
	"go_pro/internal/database"
	"go_pro/internal/loggers"
	"go_pro/internal/mailers"
	"go_pro/internal/repositories"
	"go_pro/internal/router"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
)

func main() {
	config := config.LoadConfig()
	log := loggers.NewLogger(os.Stderr, config.GetLevelLog())
	port := fmt.Sprintf(":%s", config.ServerPort)
	db, err := database.LoadDataBase(config.DBConnURL)

	sessionManager := scs.New()
	sessionManager.Lifetime = time.Hour
	sessionManager.Store = pgxstore.New(db)

	pgxstore.NewWithCleanupInterval(db, 30*time.Minute)

	if err != nil {
		slog.Error("Failed connection db", "error", err)
		panic("Server Error!")
	}

	mailService := mailers.NewSMTPMailService(mailers.SMTPConfig{
		Host:     "localhost",
		Port:     1025,
		Username: "",
		Password: "",
		From:     "quicknotes@gmail.com",
	})

	mailService.Send(mailers.MailMessage{
		To:      []string{"can@gmail.com"},
		Subject: "O Sabará Saberá!",
		Body:    []byte("Esta é uma pagina"),
	})

	noteRepo := repositories.NewNoteRepository(db)
	userRepo := repositories.NewUserRepository(db)

	slog.SetDefault(log)
	slog.Info(fmt.Sprintf("Servidor rodando na porta %s\n", config.ServerPort))

	mux := router.LoadRoutes(sessionManager, db, noteRepo, userRepo)

	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))

	if err = http.ListenAndServe(port, sessionManager.LoadAndSave(csrfMiddleware(mux))); err != nil {
		slog.Error("Server Error", "error", err)
		panic("Server Error!")
	}
}

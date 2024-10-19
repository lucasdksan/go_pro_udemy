package main

import (
	"fmt"
	"go_pro/config"
	"go_pro/internal/database"
	"go_pro/internal/handlers"
	"go_pro/internal/loggers"
	"go_pro/internal/repositories"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	config := config.LoadConfig()
	log := loggers.NewLogger(os.Stderr, config.GetLevelLog())
	mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir("./assets/"))
	port := fmt.Sprintf(":%s", config.ServerPort)
	db, err := database.LoadDataBase(config.DBConnURL)

	if err != nil {
		slog.Error("Failed connection db", "error", err)
		panic("Server Error!")
	}

	noteRepo := repositories.NewNoteRepository(db)
	userRepo := repositories.NewUserRepository(db)
	noteHandlers := handlers.NewNoteHandler(noteRepo)
	userHandlers := handlers.NewUserHandler(userRepo)

	slog.SetDefault(log)
	slog.Info(fmt.Sprintf("Servidor rodando na porta %s\n", config.ServerPort))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", staticHandler))
	mux.Handle("/", handlers.HandlerWithError(noteHandlers.NoteList))
	mux.Handle("GET /notes/{id}", handlers.HandlerWithError(noteHandlers.NoteView))
	mux.Handle("GET /notes/new", handlers.HandlerWithError(noteHandlers.NoteNew))
	mux.Handle("POST /notes", handlers.HandlerWithError(noteHandlers.NoteSave))
	mux.Handle("DELETE /notes/{id}", handlers.HandlerWithError(noteHandlers.NoteDelete))
	mux.Handle("GET /notes/{id}/update", handlers.HandlerWithError(noteHandlers.NoteEdit))

	mux.Handle("GET /user/signup", handlers.HandlerWithError(userHandlers.SignupForm))
	mux.Handle("POST /user/signup", handlers.HandlerWithError(userHandlers.Signup))
	mux.Handle("GET /user/signin", handlers.HandlerWithError(userHandlers.SigninForm))
	mux.Handle("POST /user/signin", handlers.HandlerWithError(userHandlers.Signin))
	mux.Handle("GET /confirmation/{token}", handlers.HandlerWithError(userHandlers.Confirm))

	if err = http.ListenAndServe(port, mux); err != nil {
		slog.Error("Server Error", "error", err)
		panic("Server Error!")
	}
}

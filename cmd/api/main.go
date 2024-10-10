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
	noteHandlers := handlers.NewNoteHandler()
	mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir("./assets/"))
	port := fmt.Sprintf(":%s", config.ServerPort)
	db, err := database.LoadDataBase(config.DBConnURL)

	if err != nil {
		panic("Server Error!")
	}

	noteRepo := repositories.NewNoteRepository(db)

	notes, err := noteRepo.List()

	if err != nil {
		slog.Error("Failed to list notes", "error", err)
		panic("Server Error!")
	}

	fmt.Println(notes)

	slog.SetDefault(log)
	slog.Info(fmt.Sprintf("Servidor rodando na porta %s\n", config.ServerPort))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))
	mux.Handle("/", handlers.HandlerWithError(noteHandlers.NoteList))
	mux.Handle("/notes/view", handlers.HandlerWithError(noteHandlers.NoteView))
	mux.HandleFunc("/notes/create", noteHandlers.NoteCreate)

	if err = http.ListenAndServe(port, mux); err != nil {
		panic("Server Error!")
	}
}

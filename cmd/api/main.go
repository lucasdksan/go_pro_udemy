package main

import (
	"fmt"
	"go_pro/config"
	"go_pro/internal/handlers"
	"go_pro/internal/loggers"
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

	slog.SetDefault(log)
	slog.Info(fmt.Sprintf("Servidor rodando na porta %s\n", config.ServerPort))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))
	mux.HandleFunc("/", noteHandlers.NoteList)
	mux.HandleFunc("/notes/view", noteHandlers.NoteView)
	mux.HandleFunc("/notes/create", noteHandlers.NoteCreate)

	if err := http.ListenAndServe(port, mux); err != nil {
		panic("Server Error!")
	}
}

package main

import (
	"fmt"
	"go_pro/config"
	"go_pro/internal/logger"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
		"views/templates/home.html",
	}
	t, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(w, "Error executing this page", http.StatusInternalServerError)
	}

	t.ExecuteTemplate(w, "layout", nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	files := []string{
		"views/components/footer.html",
		"views/components/header.html",
		"views/components/layout.html",
		"views/templates/note-view.html",
	}

	if id == "" {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(w, "Error executing this page", http.StatusInternalServerError)
	}

	t.ExecuteTemplate(w, "layout", id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criar uma Nota")
}

func main() {
	config := config.LoadConfig()
	log := logger.NewLogger(os.Stderr, config.GetLevelLog())

	slog.SetDefault(log)

	slog.Info(fmt.Sprintf("Servidor rodando na porta %s\n", config.ServerPort))

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/notes/view", noteView)
	mux.HandleFunc("/notes/create", noteCreate)

	port := fmt.Sprintf(":%s", config.ServerPort)

	if err := http.ListenAndServe(port, mux); err != nil {
		panic("Server Error!")
	}
}

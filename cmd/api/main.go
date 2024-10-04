package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	fmt.Println("Servidor rodando na porta 5000")

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/notes/view", noteView)
	mux.HandleFunc("/notes/create", noteCreate)

	if err := http.ListenAndServe(":5050", mux); err != nil {
		panic("Server Error!")
	}
}

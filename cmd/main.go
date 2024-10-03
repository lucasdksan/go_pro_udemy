package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Listagem de Notas")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Visualizar uma Nota")
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

	mux.HandleFunc("/notes", noteList)
	mux.HandleFunc("/notes/view", noteView)
	mux.HandleFunc("/notes/create", noteCreate)

	if err := http.ListenAndServe(":5050", mux); err != nil {
		panic("Server Error!")
	}
}
package examples

import (
	"fmt"
	"net/http"
)

func HelloHandle_6(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Oi delicia"))
}

type MyHandler_6 struct{}

func (MyHandler_6) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func Exec_6() {
	fmt.Println("Servidor rodando na porta 5000")
	m := MyHandler_6{}
	mux := http.NewServeMux()
	mux.Handle("/", m)
	mux.HandleFunc("/ioi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oi delicia oiii 12"))
	})

	mux.HandleFunc("/ioi/delicia", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oi delicia oiii 12 deliciaaaaaaaaaa"))
	})

	http.ListenAndServe(":5000", mux)
}

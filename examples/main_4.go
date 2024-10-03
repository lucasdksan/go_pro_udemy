package examples

import "net/http"

type WorldHandler struct{}

func (WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

type HelloHandler struct{}

func (HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func Exec_4() {
	handler := WorldHandler{}
	hello := HelloHandler{}

	http.Handle("/", handler)
	http.Handle("/hello", hello)

	http.ListenAndServe(":5000", nil)
}

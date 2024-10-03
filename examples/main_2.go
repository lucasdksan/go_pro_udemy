package examples

import "net/http"

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func Exec_2() {
	handler := MyHandler{}

	http.ListenAndServe(":5000", handler)
}

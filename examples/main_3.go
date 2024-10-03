package examples

import "net/http"

type MyHandler_1 struct{}

func (MyHandler_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func Exec_3() {
	handler := MyHandler_1{}

	http.Handle("/", handler)

	http.ListenAndServe(":5000", nil)
}

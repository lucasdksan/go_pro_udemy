package examples

import (
	"net/http"
)

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Oi delicia"))
}

func Exec_5() {
	http.HandleFunc("/hello", HelloHandle)

	http.ListenAndServe(":5000", nil)
}

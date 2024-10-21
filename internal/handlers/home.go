package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "home.html", nil, http.StatusOK)
}

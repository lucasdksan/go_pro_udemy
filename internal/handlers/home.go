package handlers

import (
	"go_pro/internal/render"
	"net/http"
)

type homeHandler struct {
	render *render.RenderTemplate
}

func NewHomeHandler(render *render.RenderTemplate) *homeHandler {
	return &homeHandler{render: render}
}

func (hh *homeHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	hh.render.RenderPage(w, r, "home.html", nil, http.StatusOK)
}

package handlers

import (
	"errors"
	"go_pro/internal/apperrors"
	"log/slog"
	"net/http"
	"text/template"
)

var ErrorBadRequest = func(text string) error {
	return apperrors.NewWithStatus(errors.New(text), http.StatusBadRequest)
}

var ErrorInternalServer = func(text string) error {
	return apperrors.NewWithStatus(errors.New(text), http.StatusInternalServerError)
}

var ErrorNotFound = func(text string) error {
	return apperrors.NewWithStatus(errors.New(text), http.StatusNotFound)
}

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func (f HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var statusError apperrors.StatusError
		var repoError apperrors.RepositoryError

		if errors.As(err, &statusError) {
			if statusError.HTTPStatus() == http.StatusNotFound {
				files := []string{
					"views/components/footer.html",
					"views/components/header.html",
					"views/components/layout.html",
					"views/templates/404.html",
				}

				t, err := template.ParseFiles(files...)

				if err != nil {
					http.Error(w, err.Error(), statusError.HTTPStatus())
				}

				t.ExecuteTemplate(w, "layout", statusError.Error())

				return
			}

			http.Error(w, err.Error(), statusError.HTTPStatus())
			return
		}

		if errors.As(err, &repoError) {
			slog.Error(err.Error())
			http.Error(w, "an error occurred while executing this operation", http.StatusInternalServerError)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

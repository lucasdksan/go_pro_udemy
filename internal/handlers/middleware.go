package handlers

import (
	"errors"
	"go_pro/internal/apperrors"
	"go_pro/internal/render"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type authMiddleware struct {
	session *scs.SessionManager
}

type errorHandlerMiddleware struct {
	render *render.RenderTemplate
}

func NewErrorHandlerMiddleware(render *render.RenderTemplate) *errorHandlerMiddleware {
	return &errorHandlerMiddleware{render: render}
}

func NewAuthMiddleware(session *scs.SessionManager) *authMiddleware {
	return &authMiddleware{session: session}
}

func (ah *authMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := ah.session.GetInt64(r.Context(), "userId")

		if userId == 0 {
			slog.Warn("usuário não está logado")
			http.Redirect(w, r, "/user/signin", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (ehm *errorHandlerMiddleware) HandlerError(next func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			var statusError apperrors.StatusError
			var repoError apperrors.RepositoryError

			if errors.As(err, &statusError) {
				if statusError.HTTPStatus() == http.StatusNotFound {
					slog.Error(err.Error())
					ehm.render.RenderPage(w, r, "404.html", nil, http.StatusNotFound)
					return
				}
			}

			if errors.As(err, &repoError) {
				slog.Error(err.Error())
				ehm.render.RenderPage(w, r, "generic-error.html", "an error occurred while executing this operation", http.StatusInternalServerError)
				return
			}

			slog.Error(err.Error())
			ehm.render.RenderPage(w, r, "generic-error.html", err.Error(), http.StatusInternalServerError)
			return
		}

	})
}

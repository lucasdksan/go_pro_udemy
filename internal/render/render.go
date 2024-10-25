package render

import (
	"bytes"
	"go_pro/internal/apperrors"
	"go_pro/views"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/csrf"
)

type RenderTemplate struct {
	session *scs.SessionManager
}

func NewRender(session *scs.SessionManager) *RenderTemplate {
	return &RenderTemplate{session: session}
}

func getTemplatePageFiles(t *template.Template, page string) (*template.Template, error) {
	return t.ParseFS(views.Files, "components/footer.html", "components/header.html", "components/layout.html", "templates/"+page)
}

func (rt *RenderTemplate) RenderPage(w http.ResponseWriter, r *http.Request, page string, data interface{}, status int) error {
	// files := []string{
	// 	"views/components/footer.html",
	// 	"views/components/header.html",
	// 	"views/components/layout.html",
	// }

	// files = append(files, fmt.Sprintf("views/templates/%s", page))
	t := template.New("").Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
		"csrfToken": func() string {
			return csrf.Token(r)
		},
		"isAuthenticated": func() bool {
			return rt.session.Exists(r.Context(), "userId")
		},
		"userEmail": func() string {
			return rt.session.GetString(r.Context(), "userEmail")
		},
	})

	t, err := getTemplatePageFiles(t, page)

	if err != nil {
		return apperrors.ErrorInternalServer("error executing this page")
	}

	buff := &bytes.Buffer{}

	if err = t.ExecuteTemplate(buff, "layout", data); err != nil {
		return apperrors.ErrorInternalServer("error in template")
	}

	w.WriteHeader(status)
	buff.WriteTo(w)

	return nil
}

func (rt *RenderTemplate) RenderMailBody(mailTemplate string, data interface{}) ([]byte, error) {
	t, err := template.ParseFiles("views/mails/" + mailTemplate)

	if err != nil {
		slog.Error("Error template mails generate", "error ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}

	if err := t.Execute(buffer, data); err != nil {
		slog.Error("Error template mails generate", "error ", err.Error())
		return nil, err
	}

	return buffer.Bytes(), nil
}

package http

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type TemplateRegistry struct {
	Templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func RegistTemplate() *TemplateRegistry {
	return &TemplateRegistry{
		Templates: template.Must(template.ParseGlob("internal/http/static/template/*.html")),
	}
}

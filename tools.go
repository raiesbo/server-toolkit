package servertools

import (
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
)

type Tools struct {
	TmplsDir    string
	BaseTmplDir string
}

// RenderTmpl handles the rendering of templates including the main page template and the base layout
func (t *Tools) RenderTmpl(w http.ResponseWriter, r *http.Request, name string, code int, data any) error {
	if &t.TmplsDir == nil {
		return errors.New("the templates directory (TmplsDir) variably needs to be set")
	}

	tmpl := template.New(name)

	if &t.BaseTmplDir != nil {
		tmplNew, err := tmpl.ParseFiles(t.BaseTmplDir)
		if err != nil {
			return err
		}
		tmpl = tmplNew
	}

	tmpl, err := tmpl.ParseFiles(filepath.Join(t.TmplsDir, name))
	if err != nil {
		return err
	}

	if err = tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		return err
	}

	return nil
}

// Package view provides ...
package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layouts"

type View struct {
	Template *template.Template
	Layout   string
}

func (view *View) Render(w http.ResponseWriter, data interface{}) error {
	return view.Template.ExecuteTemplate(w, view.Layout, data)
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}

	return files
}

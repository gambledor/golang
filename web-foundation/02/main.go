// Package main provides ...
package main

import (
	"github.com/gambledor/golang/web-foundation/02/views"
	"log"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layouts"

// var index *template.Template
// var contact *template.Template
var index *views.View
var contact *views.View

func check(err error, w http.ResponseWriter) {
	if err != nil {
		log.Fatal(err)
		if w != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			panic(err)
		}
	}
}

func main() {
	var err error

	// files := append(layoutFiles(), "views/index.gohtml")
	// index, err = template.ParseFiles(files...)
	// check(err, nil)
	index = views.NewView("bootstrap", "views/index.gohtml")
	// files = append(layoutFiles(), "views/contact.gohtml")
	// contact, err = template.ParseFiles(files...)
	// check(err, nil)
	contact = views.NewView("bootstrap", "views/contact.gohtml")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
	log.Fatal("ListenAndServe:", err)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("enter handler " + r.Method)
	w.Header().Set("Content-type", "text/html")
	// err := index.ExecuteTemplate(w, "bootstrap", nil)
	// check(err, w)
	index.Render(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("enter contactHandler")
	// err := contact.ExecuteTemplate(w, "bootstrap", nil)
	// check(err, w)
	contact.Render(w, nil)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	check(err, nil)

	return files
}

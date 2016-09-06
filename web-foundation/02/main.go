// Package main provides ...
package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layouts"
var index *template.Template
var contact *template.Template

func check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func main() {
	var err error

	files := append(layoutFiles(), "views/index.gohtml")
	index, err = template.ParseFiles(files...)
	check(err)
	files = append(layoutFiles(), "views/contact.gohtml")
	contact, err = template.ParseFiles(files...)
	check(err)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
	log.Fatal("ListenAndServe:", err)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("enter handler " + r.Method)
	err := index.ExecuteTemplate(w, "bootstrap", nil)
	check(err)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("enter contactHandler")
	err := contact.ExecuteTemplate(w, "bootstrap", nil)
	check(err)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	check(err)

	return files
}

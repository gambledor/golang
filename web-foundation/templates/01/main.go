//
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	wisdom := `Release self-focus; embrace other-focus.`
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", wisdom)
	if err != nil {
		log.Fatalln(err)
	}
}

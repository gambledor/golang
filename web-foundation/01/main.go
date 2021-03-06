// Package main provides setup a web server
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call it by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello globrutto!") // send data to client
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("mehod: ", r.Method) // get request method
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		checkError(w, err)
		data := struct{ Title string }{Title: "Login"}
		err = t.Execute(w, data)
		checkError(w, err)
	} else {
		r.ParseForm()
		// login part of log in
		log.Println("username: ", r.Form.Get("username"))
		log.Println("password: ", r.Form.Get("password"))
		fmt.Fprintf(w, "Hello "+r.Form.Get("username"))
	}
}

func checkError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // set Route
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

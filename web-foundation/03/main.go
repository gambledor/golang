package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	ID        int      `json:"id"`
	Moniker   string   `json:"moniker"`
	Bio       string   `json:"bio"`
	Languages []string `json"languages"`
}

// HelloWorld is a greetings type
type HelloWorld struct {
	Name string
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello "+h.Name)
}

var allUsers []*user

func init() {
	allUsers = []*user{
		{1, "Hades", "god of underworld, ruler of dead and brother to the supreme ruler of gods, Zeus", []string{"Greek"}},
		{2, "Horus", "god of the sun, sky and war", []string{"Arabic"}},
		{3, "Apollo", "god of light, music, manly beaty, dance, prophecy, medicine, poetry and almost every other thing. Son of Zeus", []string{"Greek"}},
		{4, "Artemis", "goddes of wilderness and wild animals, Sister of Apollo and daughter of Zeus", []string{"Greek"}},
	}
}

func (u user) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("method: " + r.Method)
	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		s := r.URL.Path[len("/users"):]
		if s != "" {
			id, _ := strconv.Atoi(s)
			var (
				requestedUser *user
				found         bool
			)

			for _, v := range allUsers && !found {
				if v.ID == id {
					found = true
					requestedUser = v
					break
				}
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(allUsers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(j))
}

func main() {
	// associate url requests to functions that handle requests
	http.Handle("/hello/larne", &HelloWorld{"Larne"})
	http.Handle("/hello/doe", &HelloWorld{"John doe"})
	http.Handle("/users/", user{})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "you just reached the page of our startup. Thank you for visiting")
	})

	// start web server
	log.Fatal(http.ListenAndServe(":9999", nil))
}

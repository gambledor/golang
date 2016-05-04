package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Payload type
type Payload struct {
	Stuff Data
}

// Data type
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

// Fruits type
type Fruits map[string]int

// Vegetables type
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	resp, err := getJSONResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(resp))
}

func getJSONResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	vegetables := make(map[string]int)
	vegetables["Carots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}

func main() {
	http.HandleFunc("/", serveRest)
	log.Fatal(http.ListenAndServe(":1337", nil))
}

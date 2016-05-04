package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	url := "http://localhost:1337"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var p Payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Stuff.Fruit, "\n", p.Stuff.Veggies)
}

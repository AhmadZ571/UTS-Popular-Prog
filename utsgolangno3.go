package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
}

func main() {

	url := "http://localhost:8000/"

	data1 := Person{"Ahmad Zaini"}
	data2 := Person{"Farhan Irwan"}

	b1, err := json.Marshal(data1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resp1, err := http.Post(url, "application/json", bytes.NewBuffer(b1))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp1.Body.Close()

	b2, err := json.Marshal(data2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resp2, err := http.Post(url, "application/json", bytes.NewBuffer(b2))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp2.Body.Close()

	resp3, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp3.Body.Close()

	var people []Person
	err = json.NewDecoder(resp3.Body).Decode(&people)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, person := range people {
		fmt.Println(person.Name)
	}
}

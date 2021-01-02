package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Info   string `json:"info"`
}

func main() {
	book := Book{Title: "Learn Go with joy", Author: "Nicola Solazzo", Info: "this book is Free"}
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}

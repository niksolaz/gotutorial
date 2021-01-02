package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
	Info   Info   `json:"info"`
}

type Author struct {
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
	Skill     string `json:"skill"`
}

type Info struct {
	Price       float32 `json:"price"`
	Pages       int     `json:"pages"`
	Description string  `json:"description"`
}

func main() {
	author := Author{Nickname: "Nicola Solazzo", Age: 44, Developer: true, Skill: "Golang"}
	info := Info{Price: 29.99, Pages: 320, Description: "This book is un fake integration to try your skill with Golang"}
	book := Book{Title: "Learn Go with joy", Author: author, Info: info}

	byteArray, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}

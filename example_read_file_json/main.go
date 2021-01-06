package main

import (
	"fmt"
	"os"
)

func main() {
	jsonfile, err := os.Open("info.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json file opened")

	defer jsonfile.Close()
}

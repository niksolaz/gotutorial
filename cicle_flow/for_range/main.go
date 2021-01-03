package main

import "fmt"

func main() {
	s := "This is my string range for cicle for"
	for i, v := range s {
		fmt.Println(i, v)
	}
}

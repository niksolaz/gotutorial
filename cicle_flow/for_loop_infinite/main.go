package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 9 {
			break
		}
		fmt.Println("loop infinito")
		i++
	}
}

package main

import (
	"fmt"
)

func frase(x string) string {
	fmt.Println(x)
	return x
}

func negativo(x int) bool {
	defer fmt.Println("Stack negativo()")
	if x < 10 {
		fmt.Println("Negativo")
		return true
	}
	fmt.Println("Positivo")
	return false
}

func main() {
	defer fmt.Println("Stack main()")
	negativo(-10)
	defer frase("Stack frase()")
}

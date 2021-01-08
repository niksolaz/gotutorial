package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	fmt.Println("Start Counter")
	for i := 10; i > value; i-- {
		time.Sleep(time.Second)
		fmt.Println(i)
		if i == 0 {
			fmt.Println("Start Shuttle!!!")
		}
	}
}

func main() {
	fmt.Println("Goroutines")

	go compute(-1)
	go compute(-1)

	var input string
	fmt.Scanln(&input)
}

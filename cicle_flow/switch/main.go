package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digit your age in the Shell")
	fmt.Println("---------------------")
	if scanner.Scan() {
		myAge := scanner.Text()
		fmt.Println("---------------------")
		b := digitCheck.MatchString(myAge)
		if b {

			switch {
			case myAge <= "40":
				fmt.Println("you are young man")
			case myAge > "40":
				fmt.Println("you are man")
			}
		} else {
			fmt.Println("This is not an age validate")
		}
	}
}
